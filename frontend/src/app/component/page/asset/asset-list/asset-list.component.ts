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
import { NzPaginationModule } from 'ng-zorro-antd/pagination'
import { Router } from '@angular/router'
import { NzSelectModule } from 'ng-zorro-antd/select'
import { QRcodeComponent } from '../../../components/qr-code/qr-code.component'
import { RepairRecordCreateComponent } from '../repair-record-create/repair-record-create.component'
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker'
import { UserStoreService } from '../../../../../state/user.service'
import { findMenuItem } from '../../../tool-function'
import { Subscription } from 'rxjs'
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
        QRcodeComponent,
        RepairRecordCreateComponent,
        NzDatePickerModule,
        DownloadExcelTemplateComponent,
        UploadDialogComponent
    ],
    templateUrl: './asset-list.component.html',
    styleUrl: './asset-list.component.css',
})
export class AssetListComponent {
    private rightSubscription: Subscription
    constructor(
        private routeTo: Router,
        private userStoreService: UserStoreService
    ) {
        this.rightSubscription = this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Tax Information', 'tax-information')
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

    searchForm: any = {
        page: 1,
        limit: 10
    }

    userRightInside: any = {
        read: false,
        write: false,
        update: false,
        delete: false
    }


    dataLists: any[] = []
    totals: number = 0
    editFormDialog: boolean = false
    removeDialog: boolean = false
    handleId: string = ''

    ngOnInit() {
        this.loadAssetListLists()
        this.loadTypeList()
        this.loadDeptList()
        this.loadLocationList()
    }

    typeLists: any[] = []
    async loadTypeList() {
        this.typeLists = await getApiWithAuth('/asset/type/getAll')
    }

    deptLists: any[] = []
    async loadDeptList() {
        this.deptLists = await getApiWithAuth('/sys/department/getAll')
    }

    placeLists: any[] = []
    async loadLocationList() {
        this.placeLists = await getApiWithAuth('/base/location/getAll')
    }

    async loadAssetListLists() {
        const res = await postApiWithAuth('/asset/asset-list/list', this.searchForm)
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

    openEdit(id: string) {
        this.routeTo.navigate([`/asset-update`], { queryParams: { id } })
    }

    goToCreate() {
        this.routeTo.navigate(['/asset-create'])
    }

    goToWriteOff() {
        this.routeTo.navigate([`write-off`], { queryParams: { id: this.handleId }})
    }

    qrCodeDialog: boolean = false
    qrCodeString: string = ''
    qrCodeRandomHtml: string = ''

    openQrCodeDialog(data: any) {
        this.qrCodeDialog = true

        const qrCodeString = `${data.assetCode}|${data.assetName}|${this.dateFormat(data.purchaseDate)}|${data.assettype.typeCode}|${data.assettype.typeName}|${data.location.placeCode}|${data.location.placeName}|${data.department.deptCode}|${data.department.deptName}`
        this.qrCodeString = qrCodeString

        this.qrCodeRandomHtml = `
            <div class="px-3 text-left grid grid-cols-1">
                <div><span class="font-bold">Asset Code:</span>${data.assetCode}</div>
                <div><span class="font-bold">Asset Name:</span> ${data.assetName}</div>
                <div><span class="font-bold">Purchase Date:</span> ${this.dateFormat(data.purchaseDate)}</div>
                <div><span class="font-bold">Type Code:</span> ${data.assettype.typeCode}</div>
                <div><span class="font-bold">Type Name:</span> ${data.assettype.typeName}</div>
                <div><span class="font-bold">Location Code:</span> ${data.location.placeCode}</div>
                <div><span class="font-bold">Location Name:</span> ${data.location.placeName}</div>
                <div><span class="font-bold">Department Code:</span> ${data.department.deptCode}</div>
                <div><span class="font-bold">Department Code:</span> ${data.department.deptName}</div>
            </div>
        `
    }

    openQrCodeDialogClose(event: any) {
        this.qrCodeDialog = false
    }

    repairRecordDialog: boolean = false
    handleData: any = {}

    openRepairRecordDialog(data: any) {
        this.repairRecordDialog = true
        this.handleId = data._id
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