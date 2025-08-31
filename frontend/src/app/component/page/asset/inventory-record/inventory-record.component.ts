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
        NzPaginationModule,
        NzDatePickerModule,
    ],
    templateUrl: './inventory-record.component.html',
    styleUrl: './inventory-record.component.css',
})
export class InventoryRecordListComponent {
    private rightSubscription: Subscription
    constructor(
        private userStoreService: UserStoreService
    ) {
        this.rightSubscription = this.userStoreService.menuRole$.subscribe((data: any) => {
            const answer = findMenuItem(data, 'Inventory Record', 'inventory-record')
            this.userRightInside = {
                read: answer?.read ?? false
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
        read: false
    }

    searchForm: any = {
        page: 1,
        limit: 10
    }

    dataLists: any[] = []
    totals: number = 0

    ngOnInit() {
        this.loadWriteOffLists()
    }

    async loadWriteOffLists() {
        const res = await postApiWithAuth('/sys/inv-record/list', this.searchForm)
        this.dataLists = res.lists
        this.totals = res.total
    }


    dateFormat(data: string) {
        return data ? moment(new Date(data)).format('DD-MM-YYYY HH:MM') : null
    }


}