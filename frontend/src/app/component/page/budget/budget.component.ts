import { Component, OnChanges } from '@angular/core'
import { CommonModule } from '@angular/common'
import { FormsModule } from '@angular/forms'
import { deleteApiWithAuth, getApiWithAuth, postApiWithAuth } from '../../../../tool/httpRequest-auth'
import { NzTableModule } from 'ng-zorro-antd/table'
import { NzButtonModule } from 'ng-zorro-antd/button'
import { NzModalModule, NzModalService } from 'ng-zorro-antd/modal'
import { NzInputModule } from 'ng-zorro-antd/input'
import { NzFormModule } from 'ng-zorro-antd/form'
import moment from 'moment'
import { BudgetForm } from './interface'
import { NzMessageService } from 'ng-zorro-antd/message'
import { NzPaginationModule } from 'ng-zorro-antd/pagination'
import { NzSelectModule } from 'ng-zorro-antd/select'
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker'
import { NzInputNumberModule } from 'ng-zorro-antd/input-number'
import { findMenuItem } from '../../tool-function'
import { UserStoreService } from '../../../../state/user.service'
import { Subscription } from 'rxjs'
import { downloadTempExcelFile } from '../../../../tool/excel-helper'
import { DownloadExcelTemplateComponent } from '../../components/download-template-component/download-template-component.component'
import { UploadDialogComponent } from '../../components/upload-dialog-component/upload-dialog-component.component'
import { DownloadExcelDataComponent } from '../../components/download-excel-component/download-excel-data-component.component'
import { FilterListApiSmall, ListOption } from '../../components/filter-list-small/filter-list.component'

@Component({
    // selector: 'app-footer',
    standalone: true,
    imports: [
        CommonModule, 
        NzFormModule, 
        NzButtonModule, 
        FormsModule, 
        NzModalModule, 
        NzTableModule, 
        NzInputModule, 
        NzPaginationModule,
        NzSelectModule,
        NzDatePickerModule,
        NzInputNumberModule,
        DownloadExcelTemplateComponent,
        UploadDialogComponent,
        DownloadExcelDataComponent,
        FilterListApiSmall
    ],
    templateUrl: './budget.component.html',
    styleUrl: './budget.component.css',
})
export class BudgetComponent {
    private rightSubscription: Subscription
    constructor(
        private message: NzMessageService,
        private userStoreService: UserStoreService
    ) {
        this.rightSubscription = this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Budget', 'budget')
            console.log(answer)
            this.userRightInside = {
                read: answer?.read ?? false,
                write: answer.write ?? false,
                update: answer.update ?? false,
                delete: answer.delete ?? false,
                upload: answer.upload ?? false
                 // keep default value
            }
            this.excelFileSetting.code = answer?.excelFunctionCode ?? ''
            this.preLoadExcelSetting()
        })
                
    }

    ngOnDestroy() {
        if (this.userStoreService.menuRole$) {
            this.rightSubscription.unsubscribe()
        }
    }

    userRightInside: any = {
        read: false,
        write: false,
        update: false,
        delete: false,
        upload: false,
    }

    searchForm: any = {
        date: [],
        deptIds: [],
        placeIds: [],
        page: 1,
        limit: 10
    }

    editForm: BudgetForm =  {
        id: '',
        deptId: '',
        placeId: '',
        budgetNo: '',
        budgetName: '',
        year: '',
        month: '',
        budgetAmount: 0,
        budgetFrom: '',
        budgetTo: '',
        budgetStatus: '',
        remark: '',
    }

    okText: string = 'Create'

    dataLists: any[] = []
    totals: number = 0
    editFormDialog: boolean = false
    removeDialog: boolean = false
    handleRemoveId: string = ''
    taxInformation: boolean = false

    fileList: any[] = []
    deptLists: any[] = []
    placeLists: any[] = []

    department?: any = {}


    ngOnInit() {
        this.loadBudgetLists()
        this.loadDeptList()
        this.loadLocationList()
    }

    async loadDeptList() {
        const res = await getApiWithAuth('/sys/department/all')
        this.deptLists = res.data
    }

    async loadLocationList() {
        const res = await getApiWithAuth('/base/location/all')
        this.placeLists = res.data
    }


    async submitForm() {
        const url = this.editForm.id === '' ? '/base/budget/create' : `/base/budget/update`

        console.log(this.editForm)
        const res = await postApiWithAuth(url, this.editForm)

        if (res.error) {
            this.message.error(res.error)
        } else if (!res.error) {
            this.message.success('Save successful!')
            this.closeDialog()
            this.loadBudgetLists()

            this.editForm = {
                id: '',
                deptId: '',
                placeId: '',
                budgetNo: '',
                budgetName: '',
                year: '',
                month: '',
                budgetAmount: 0,
                budgetFrom: '',
                budgetTo: '',
                budgetStatus: '',
                remark: '',
            }

            this.okText = 'Create'
        }
    }

    async loadBudgetLists() {
        const res = await postApiWithAuth('/base/budget/list', this.searchForm)
        this.dataLists = res.lists
        this.totals = res.total
    }

    showDialog() {
        this.editFormDialog = true
    }

    closeDialog() {
        this.editFormDialog = false
    }

    handleRomeve(id: string) {
        this.handleRemoveId = id
        this.removeDialog = true
    }

    closeRemoveDialog() {
        this.removeDialog = false
    }

    async handleRemove() {
        const url = `/base/budget/void/${this.handleRemoveId}`

        const res: any = await deleteApiWithAuth(url)

        if (res.error) {
            this.message.error(res.error)
        } else if (!res.error) {
            this.message.success('Delete successful!')
            this.closeRemoveDialog()
            this.loadBudgetLists()
        }
    }

    dateFormat(data: string) {
        return data ? moment(data).format('DD-MM-YYYY HH:mm') : null
    }

    async getOneData(id: string) {
        const res = await getApiWithAuth(`/base/budget/one/${id}`)
        if (res.error) {
            this.message.error(res.error)
        } else if (!res.error) {
            this.editForm = res
            this.okText = 'Update'
            this.showDialog()
        }
    }

    excelFileSetting: any = {
        code: ''
    }

    dbFieldList: string[] = []
    excelFieldList: string[] = []
    async preLoadExcelSetting() {
        const res = await getApiWithAuth(`/sys/excel-field-match/code/${this.excelFileSetting.code}`)
        this.dbFieldList = res.fieldLists.map((item: any) => item.dbFieldName)
        this.excelFieldList = res.fieldLists.map((item: any) => item.excelFieldName)
    }
    // Department data
    deptTotal: number = 0
    depts: any = []
    deptOptions: ListOption[] = [];
    deptSearchForm: any = {
        deptCode: '',
        deptName: '',
        page: 1,
        limit: 10
    }
    loadDeptLists = async (
        url: string,
        body: Record<string, unknown>,
    ): Promise<ListOption[]> => {
        const payload = {
          ...this.deptSearchForm,
          ...body,
        }
    
        try {
            const res = await postApiWithAuth(url, payload)
        
            this.depts = res.data.lists ?? []
            this.deptTotal = res.data.total ?? 0
        
            this.deptOptions = this.depts.map((a: any) => ({
                value: a.id,
                label: `${a.deptCode}`,
            }));
        
            return this.deptOptions
        } catch (error) {
            this.message.error('API request failed')
            throw error
        }
    }

    // Location data
    placeTotal: number = 0
    places: any = []
    placeOptions: ListOption[] = [];
    placeSearchForm: any = {
        placeCode: '',
        placeName: '',
        page: 1,
        limit: 10
    }
    loadPlaceLists = async (
        url: string,
        body: Record<string, unknown>,
    ): Promise<ListOption[]> => {
        const payload = {
          ...this.placeSearchForm,
          ...body,
        }
    
        try {
            const res = await postApiWithAuth(url, payload);
        
            this.places = res ?? [];
            this.placeTotal = res.length ?? 0;
        
            this.placeOptions = this.places.map((a: any) => ({
                value: a.id,
                label: `${a.placeCode}`,
            }))
    
            return this.placeOptions
        } catch (error) {
            this.message.error('API request failed')
            throw error
        }
    }
}