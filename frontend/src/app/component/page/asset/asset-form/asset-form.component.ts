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
import { debounceTime, Observable, Observer, Subject, Subscription, timer } from 'rxjs'
import { MatButtonModule } from '@angular/material/button'
import { MatIconModule } from '@angular/material/icon'
import { NzUploadChangeParam, NzUploadFile, NzUploadModule } from 'ng-zorro-antd/upload'
import { NzIconModule } from 'ng-zorro-antd/icon'
import { uploadImgToBase64 } from '../../../../../tool/imageUpload'
import { FileViewComponent } from '../../../components/file-view-dialog/file-view-dialog.component'
import { findMenuItem } from '../../../tool-function'
import { UserStoreService } from '../../../../../state/user.service'
// import { UploadComponentComponent } from '../../../components/upload-component/upload-component.component'

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
        MatButtonModule, 
        NzUploadModule,
        NzButtonModule, 
        NzIconModule,
        NzModalModule,
        FileViewComponent
    ],
    templateUrl: './asset-form.component.html',
    styleUrl: './asset-form.component.css',
})
export class AssetFormComponent implements OnInit {
    private rightSubscription: Subscription
    constructor(
        private route: ActivatedRoute, 
        private routeTo: Router,
        private message: NzMessageService,
        private userStoreService: UserStoreService
    ) {
        this.changeEvent$.pipe(debounceTime(300)).subscribe(event => {
            this.preAction(event.file.originFileObj);
        })
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
        update: false
    }

    private changeEvent$ = new Subject<NzUploadChangeParam>()

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
        uploadAssetListFiles: []
    }

    taxInformation: boolean = false

    theId: any = ''

    fileUpdloadList: any = []


    ngOnInit() {
        this.route.queryParams.subscribe((x: any) => {
            if (x.id) {
                this.theId = x.id
                this.getOne()
            }
        })


      /*  if (this.route.snapshot.paramMap.get('id')) {
            this.theId = this.route.snapshot.paramMap.get('id')
            this.getOne()
        } */

        this.loadTypeList()
        this.loadDeptList()
        this.loadLocationList()
        this.loadVendorList()
        this.loadTaxInfoList()
    }

    beforeUpload = (file: NzUploadFile, _fileList: NzUploadFile[]): Observable<boolean> =>
        new Observable((observer: Observer<boolean>) => {
        const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
        if (!isJpgOrPng) {
            this.message.error('You can only upload JPG or PNGfile!')
            observer.complete()
       //     return
        }
        const isLt2M = file.size! / 1024 / 1024 < 3
        if (!isLt2M) {
                this.message.error('Image must smaller than 3MB!')
             //   observer.complete()
  //      return
        }
    
        observer.next(isJpgOrPng && isLt2M)
        observer.complete()
    })

    finalFileList: any[] = []
    handleChange(event: any) {
        switch(event.type) {
            case 'start':
                this.changeEvent$.next(event)
            break

            case 'removed':
                this.finalFileList = this.finalFileList.filter(f => f.fileName !== event.file.originFileObj.name)
            break
        }
        console.log(event.type)
        console.log(this.finalFileList)
        
    }

    async preAction(fileObj: any) {
        const response: any = await uploadImgToBase64(fileObj)
        this.finalFileList.push({
            fileName: fileObj.name,
            fileType: fileObj.type,
            base64: response.data
        })
    }

    async getOne() {
        this.editForm = await getApiWithAuth(`/asset/asset-list/one/${ this.theId}`)
        console.log(this.editForm, 'k')
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

    vendorLists: any[] = []
    async loadVendorList() {
        this.vendorLists = await getApiWithAuth('/base/vendor/getAll')
    }

    taxLists: any[] = []
    async loadTaxInfoList() {
        const results = await getApiWithAuth('/base/tax-information/getAll')
        const updates = results.map((x: any)=> {
            return {
              ...x,
              taxRate: Number(x.taxRate),
              label: `${x.countryName} ${x.taxCode} (${Number(x.taxRate) * 100}%)`
            }
        })

        this.taxLists = updates
    }
    
    taxInfoOnChanges(event: Event) {
        const data = this.taxLists.find((item: any) => item._id === event)
        this.editForm.taxInfofId = data._id
        this.editForm.taxCountryCode = data.countryCode
        this.editForm.taxCode = data.taxCode
        this.editForm.taxRate = data.taxRate
    }

    taxOnChanges(event: any) {
        const cost: number = this.editForm.cost 
        const taxRate: number = this.editForm.taxRate

        switch(event) {
            case true:
                this.editForm.afterBeforeTax = cost * (1 - taxRate)
            break
            case false:
                this.editForm.afterBeforeTax = cost * (1 + taxRate)
            break
        }
    }

    async submitForm() {
        const url = this.editForm._id ? '/asset/asset-list/update' : '/asset/asset-list/create'
        this.editForm.uploadAssetListFiles = this.finalFileList

        const res = await postApiWithAuth(url, this.editForm)

        if (!res.msg) {
            this.message.info('Data save successfully!')
            timer(2500).subscribe(() => {
                this.routeTo.navigate(['/asset-list'])
            })
        }
    }

    resetForm() {
        this.editForm = {
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
            uploadAssetListFiles: []
        }
    }

    backToList() {
        this.routeTo.navigate([`/asset-lists`])
    }

    fileDialogVisible: boolean = false
    openFileDialog() {
        this.fileDialogVisible = true
    }

    closeFileDialog() {
        this.fileDialogVisible = false
    }

}