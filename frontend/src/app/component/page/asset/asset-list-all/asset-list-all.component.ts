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
import { UserStoreService } from '../../../../../state/user.service'
import { findMenuItem } from '../../../tool-function'
import { Subscription } from 'rxjs'

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
        NzPaginationModule
    ],
    templateUrl: './asset-list-all.component.html',
    styleUrl: './asset-list-all.component.css',
})
export class AssetListAllComponent {
    private rightSubscription: Subscription
    constructor(
        private userStoreService: UserStoreService,
        private routeTo: Router
    ) {
        this.rightSubscription = this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Asset List Report', 'asset-list-all')
            this.userRightInside = {
                read: answer?.read ?? false
                 // keep default value
            }
        })
        this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Asset List Report', 'asset-list-all')
                                                            
            this.userRightInside = {
                read: answer.read
            }
        })
    }

    ngOnDestroy() {
        if (this.userStoreService.menuRole$) {
            this.rightSubscription.unsubscribe()
        }
    }

    userRightInside: any = {
        read: false
    }

    searchForm: any = {
        page: 1,
        limit: 10
    }

    okText: string = 'Create'

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
        this.dataLists = await getApiWithAuth('/asset/asset-list/list-all')
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
        this.routeTo.navigate([`/asset-update/${id}`])
    }

    goToCreate() {
        this.routeTo.navigate(['/asset-create'])
    }

    goToWriteOff() {
        this.routeTo.navigate([`write-off/${this.handleId}`])
    }

}