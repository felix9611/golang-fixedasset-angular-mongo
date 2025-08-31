import { Component, EventEmitter, Input, OnInit, Output, SimpleChanges } from '@angular/core'
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
import { timer } from 'rxjs'

@Component({
    selector: 'app-repair-record-create',
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
        NzInputNumberModule
    ],
    templateUrl: './repair-record-create.component.html',
    styleUrl: './repair-record-create.component.css',
})
export class RepairRecordCreateComponent implements OnInit {
    @Input() repairRecordDialog!: boolean
    @Input() handleId!: string
    @Output() repairRecordDialogChange = new EventEmitter<boolean>()

    constructor(
        private route: ActivatedRoute, 
        private routeTo: Router,
        private message: NzMessageService
    ) {}

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
        repairReason: '',
        maintenanceReriod: false,
        maintenanceName: '',
        maintenanceDate: '',
        maintenanceFinishDate: '',
        repairInvoiceDate: '',
        repairInvoiceNo: '',
        repairAmount: 0,
        remarkRepair: '',
    }

    ngOnChanges(changes: SimpleChanges) {
        this.getOne(changes['handleId'].currentValue)
    }

    ngOnInit() {
        
    }

    async getOne(id: string) {
       this.editForm = await getApiWithAuth(`/asset/asset-list/one/${ id}`)
        
    }


    backTo() {
        this.routeTo.navigate(['/asset-list'])
    }

    removeDialog: boolean = false


    async goToSave() {
        const finalForm = {
            assetId: this.editForm._id,
            repairReason: this.editForm.repairReason,
            maintenanceReriod: this.editForm.maintenanceReriod,
            maintenanceName: this.editForm.maintenanceName,
            maintenanceDate: this.editForm.maintenanceDate,
            maintenanceFinishDate: this.editForm.maintenanceFinishDate,
            repairInvoiceDate: this.editForm.repairInvoiceDate,
            repairInvoiceNo: this.editForm.repairInvoiceNo,
            repairAmount: this.editForm.repairAmount,
            remark: this.editForm.remarkRepair
        }

        const res = await postApiWithAuth('/asset/repair-record/create', finalForm)
        
        if (res.finish) {
            this.message.info(res.msg)
            timer(2500).subscribe(() => {
                this.closeDialog()
            })
        } else {
            this.message.error(res.msg)
        }
    }

    closeDialog() {
        this.repairRecordDialog = false
        this.repairRecordDialogChange.emit(false)
    }
 
}