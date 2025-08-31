import { Component, OnInit } from '@angular/core'
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
import { AssetFormDto } from './interface'
import { NzSelectModule } from 'ng-zorro-antd/select'
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker'
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox'
import { NzInputNumberModule } from 'ng-zorro-antd/input-number'
import { ActivatedRoute, Router } from '@angular/router'
import { Subscription, timer } from 'rxjs'
import { MatButtonModule } from '@angular/material/button'
import { MatIconModule } from '@angular/material/icon'
import { UserStoreService } from '../../../../../state/user.service'
import { findMenuItem } from '../../../tool-function'

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
        NzCheckboxModule,
        NzInputNumberModule,
        MatIconModule, 
        MatButtonModule
    ],
    templateUrl: './write-off-form.component.html',
    styleUrl: './write-off-form.component.css',
})
export class WriteOffFormComponent implements OnInit {
    private rightSubscription: Subscription
    constructor(
        private route: ActivatedRoute, 
        private routeTo: Router,
        private message: NzMessageService,
        private userStoreService: UserStoreService
    ) {
        this.rightSubscription = this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Asset List', 'asset-lists')
            this.userRightInside = {
                read: answer?.read ?? false,
                write: answer.write ?? false,
                update: answer.update ?? false,
                delete: answer.delete ?? false,
                upload: answer.upload ?? false
                 // keep default value
            }
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
        delete: false
    }

    editForm: AssetFormDto = {
        _id: '',
        assetCode: '',
        assetName: '',
        unit: '',
        typeId: '',
        deptId: '',
        placeId: '',
        purchaseDate: '',
        description: '',
        sponsor:  false,
        sponsorName: '',
        cost: 0,
        serialNo:  '',
        invoiceNo: '',
        invoiceDate: '',
        invoiceRemark:  '',
        vendorId:  '',
        remark:  '',
        taxInfofId:  '',
        taxCountryCode:  '',
        taxCode:  '',
        taxRate:  0,
        includeTax: false,
        afterBeforeTax: 0,
        accountCode:  '',
        accountName: '',
        brandCode: '',
        brandName: '',
        chequeNo: '',
        maintenancePeriodStart: '',
        maintenancePeriodEnd: '',
        voucherNo: '',
        voucherUsedDate: '',
        assetListFiles: [],
        reason: '',
        lastDay: '',
        disposalMethod: '',
        remainingValue: 0,
        depreciationRate: 0,
        yearNumber: 0
    }

    taxInformation: boolean = false

    theId: any = ''


    ngOnInit() {
        this.route.queryParams.subscribe((x: any) => {
            if (x.id) {
                this.theId = x.id
                this.getOne()
                this.loadCodeTypeList()
                this.loadTypeList()
                this.loadDeptList()
                this.loadLocationList()
                this.loadVendorList() 
            }
        })
        
    }

    dataCal() {
        const typeData = this.typeLists.find((x: any) => x._id === this.editForm.typeId)
        this.editForm.depreciationRate = typeData.depreciationRate * 100
        const today = new Date()
        const oldDate = new Date(this.editForm.purchaseDate)

        const yearGap = today.getFullYear() - oldDate.getFullYear()
        this.editForm.yearNumber = yearGap

        const todayValue = (this.editForm.cost * typeData.depreciationRate) / yearGap
        this.editForm.remainingValue = todayValue

        
    }

    async getOne() {
        this.editForm = await getApiWithAuth(`/asset/asset-list/one/${ this.theId}`)
        
    }

    typeLists: any[] = []
    async loadTypeList() {
        this.typeLists = await getApiWithAuth('/asset/type/getAll')
        this.dataCal()
    }

    deptLists: any[] = []
    async loadDeptList() {
        this.deptLists = await getApiWithAuth('/sys/department/getAll')
    }

    placeLists: any[] = []
    async loadLocationList() {
        this.placeLists = await getApiWithAuth('/base/location/getAll')
    }

    vendorLists: any[] = []
    async loadVendorList() {
        this.vendorLists = await getApiWithAuth('/base/vendor/getAll')
    }

    typeCodeList: any[] = []
    async loadCodeTypeList() {
        this.typeCodeList = await getApiWithAuth('/base/code-type/get-type/WriteOff')
    }

    backTo() {
        this.routeTo.navigate(['/asset-list'])
    }

    removeDialog: boolean = false


    async goToWriteOff() {
        const finalForm = {
            assetId: this.editForm._id,
            lastPlaceId: this.editForm.placeId,
            reason: this.editForm.reason,
            lastDay: this.editForm.lastDay,
            disposalMethod: this.editForm.disposalMethod,
            remainingValue: this.editForm.remainingValue
        }

        const res = await postApiWithAuth('/asset/write-off/create', finalForm)
        
        if (res.finish) {
            this.message.info(res.msg)
            timer(2500).subscribe(() => {
                this.backTo()
            })
        } else {
            this.message.error(res.msg)
        }
    }

    closeRemoveDialog() {
        this.removeDialog = false
    }

    showDialog() {
        this.removeDialog = true
    }

    valueOnChange(event: any) {
        const typeData = this.typeLists.find((x: any) => x._id === this.editForm.typeId)
        const oldDate = new Date(this.editForm.purchaseDate)

        const yearGap = event.getFullYear() - oldDate.getFullYear()
        this.editForm.yearNumber = yearGap

        const todayValue = (this.editForm.cost * typeData.depreciationRate) / yearGap
        this.editForm.remainingValue = todayValue
    }

    backToList() {
        this.routeTo.navigate([`/asset-lists`])
    }
 
}