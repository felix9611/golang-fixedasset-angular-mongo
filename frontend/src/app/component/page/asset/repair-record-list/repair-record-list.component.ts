import { Component } from '@angular/core'
import { CommonModule } from '@angular/common'
import { FormsModule } from '@angular/forms'
import { getApiWithAuth, postApiWithAuth } from '../../../../../tool/httpRequest-auth'
import { NzTableModule } from 'ng-zorro-antd/table'
import { NzButtonModule } from 'ng-zorro-antd/button'
import { NzModalModule, NzModalService } from 'ng-zorro-antd/modal'
import { NzInputModule } from 'ng-zorro-antd/input'
import { NzFormModule } from 'ng-zorro-antd/form'
import moment from 'moment'
import { NzMessageService } from 'ng-zorro-antd/message'
import { NzPaginationModule } from 'ng-zorro-antd/pagination'
import { Router } from '@angular/router'
import { NzSelectModule } from 'ng-zorro-antd/select'
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker'
import { UpdateRepairRecordDto } from './interface'
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox'
import { NzInputNumberModule } from 'ng-zorro-antd/input-number'
import { Subscription, timer } from 'rxjs'
import { UserStoreService } from '../../../../../state/user.service'
import { findMenuItem } from '../../../tool-function'
import { UploadDialogComponent } from '../../../components/upload-dialog-component/upload-dialog-component.component'
import { DownloadExcelTemplateComponent } from '../../../components/download-template-component/download-template-component.component'

@Component({
    // selector: 'app-footer',
    standalone: true,
    imports: [ 
        NzSelectModule, 
        CommonModule, 
        NzFormModule, 
        NzButtonModule, 
        FormsModule, 
        NzModalModule, 
        NzTableModule, 
        NzInputModule, 
        NzPaginationModule,
        NzDatePickerModule,
        NzCheckboxModule,
        NzInputNumberModule,
        DownloadExcelTemplateComponent,
        UploadDialogComponent
    ],
    templateUrl: './repair-record-list.component.html',
    styleUrl: './repair-record-list.component.css',
})
export class RepairRecordListComponent {
    private rightSubscription: Subscription
    constructor(
        private message: NzMessageService,
        private userStoreService: UserStoreService
    ) {
        this.rightSubscription = this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Repair Record', 'repair-records')
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
        upload: false
    }

    searchForm: any = {
        page: 1,
        limit: 10
    }

    editForm: UpdateRepairRecordDto = {
        _id: '',
        assetId: '',
        repairReason: '',
        maintenanceReriod: false,
        maintenanceName: '',
        maintenanceDate: '',
        maintenanceFinishDate: '',
        repairInvoiceDate: '',
        repairInvoiceNo: '',
        repairAmount: 0,
        remark: ''
    }

    okText: string = 'Create'

    dataLists: any[] = []
    totals: number = 0
    editFormDialog: boolean = false
    removeDialog: boolean = false
    handleId: string = ''

    ngOnInit() {
        this.loadRepairRecordLists()
        this.loadLocationList()
    }

    typeLists: any[] = []


    deptLists: any[] = []


    placeLists: any[] = []
    async loadLocationList() {
        this.placeLists = await getApiWithAuth('/base/location/getAll')
    }


    async loadRepairRecordLists() {
        const res = await postApiWithAuth('/asset/repair-record/list', this.searchForm)
        this.dataLists = res.lists
        this.totals = res.total
    }

    handleRomeve(id: string) {
        this.handleId = id
        this.removeDialog = true
    }

    closeRemoveDialog() {
        this.removeDialog = false
    }


    dateFormat(data: string) {
        return data ? moment(new Date(data)).format('DD-MM-YYYY HH:MM') : null
    }

    editTitle: string = ''
    async openEdit(data: any) {
        this.editFormDialog = true
        this.editTitle = `Edit Repair Record (Code: ${data.assetlist.assetCode})`
        const id = data._id
        const getUrl = `/asset/repair-record/one/${id}`
        this.editForm = await getApiWithAuth(getUrl)
    }

    async goToSave() {
            
        const res = await postApiWithAuth('/asset/repair-record/update', this.editForm)
            
        if (res.acknowledged === true) {
            this.message.info('Update successfully!')
            this.loadRepairRecordLists()
            this.closeEditDialog()
        } else {
            this.message.error(res.msg)
        }
    }

    closeEditDialog() {
        this.editFormDialog = false
    }

    async goToWriteOff() {
        const res = await getApiWithAuth(`/asset/repair-record/void/${this.handleId}`)

        this.message.info(res.msg)

        this.closeRemoveDialog()
        this.loadRepairRecordLists()
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
}