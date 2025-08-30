import { Component, HostListener, OnInit } from '@angular/core'
import { CommonModule } from '@angular/common'
import { Router, RouterLink, RouterModule } from '@angular/router'

import { MatIconModule } from '@angular/material/icon'
import { MatButtonModule } from '@angular/material/button'
import { TooltipModule } from 'primeng/tooltip'
import path from 'path'
import { UserStoreService } from '../../../state/user.service'

@Component({
  selector: 'app-menu',
  imports: [CommonModule, RouterLink, RouterModule, MatIconModule, MatButtonModule, TooltipModule],
  templateUrl: './menu.component.html',
  styleUrl: './menu.component.css'
})
export class MenuComponent implements OnInit{
    constructor(private router: Router, private userService: UserStoreService) {
     //   let currentRoute = this.router.url.split('?')[0]

    }

    isMenuExpanded: boolean = window.innerWidth > 640 // Expand only on large screens
    screenWidth: number = window.innerWidth

    copyRights: any = {
        name: 'Felix',
        year: new Date().getFullYear()
    }

    @HostListener('window:resize', ['$event'])
    onResize(event: any) {
        this.screenWidth = window.innerWidth;
        this.isMenuExpanded = this.screenWidth > 640
    }


    ngOnInit() {
        this.userService.menu$.subscribe(menuItems => {
            this.menuItems = menuItems
        })
    }

    year: number = new Date().getFullYear()

    activeParent: any = null
    activeChild: any = null
    expandedMenus: any[] = []


    toggleMenu(item: any) {
        if (this.expandedMenus.includes(item)) {
          this.expandedMenus = this.expandedMenus.filter(menu => menu !== item) // Collapse menu
          if (this.activeParent === item) {
            this.activeParent = null
          }
        } else {
          this.expandedMenus.push(item) // Expand menu
          this.activeParent = item
        }
      }
    
    setActiveChild(child: any, parent: any) {
        this.activeParent = parent
        this.activeChild = child
    }



    menuItems: any[] = [
        {
            path: '',
            label: 'Home',
            icon: 'home'
        },
        {
            label: 'System Management',
            icon: 'settings',
            isOpen: false,
            childrens: [
                {
                    label: 'User',
                    path: 'users',
                    icon: 'account-circle'
                },
                {
                    label: 'Role',
                    path: 'role',
                    icon: ''
                },
                {
                    label: 'Menu',
                    icon: '',
                    path: 'menu'
                },
                {
                    label: 'Department',
                    path: 'department',
                    icon: 'team_dashboard'
                }
            ]
        },
        {
            label: 'Base Management',
            icon: 'dock_to_bottom',
            isOpen: false,
            childrens: [
                {
                    label: 'Tax Information',
                    icon: '',
                    path: 'tax-information'
                },
                {
                    label: 'Budget',
                    path: 'budget',
                    icon: ''
                },
                {
                    label: 'Action Log',
                    icon: '',
                    path: 'action-record'
                },
                {
                    label: 'Vendor',
                    icon: '',
                    path: 'vendor'
                },
                {
                    label: 'Location',
                    path: 'location',
                    icon: 'location'
                },
                {
                    label: 'Code Type',
                    icon: '',
                    path: 'code-type'
                }
            ]
        },
        {
            label: 'Asset Managemnt',
            icon: 'save',
            childrens: [
                {
                    label: 'Asset Type',
                    icon: '',
                    path: 'asset-type'
                },
                {
                    label: 'Asset List',
                    icon: '',
                    path: 'asset-lists'
                },
                {
                    label: 'Create Asset',
                    icon: '',
                    path: 'asset-create'
                },
                {
                    label: 'Stock Take',
                    icon: '',
                    path: 'stock-takes'
                },
                {
                    label: 'Asset List Report',
                    icon: '',
                    path: 'asset-list-all'
                },
                {
                    label: 'Inventory Record',
                    icon: '',
                    path: 'inventory-record'
                },
                {
                    label: 'Repair Record',
                    icon: '',
                    path: 'repair-records'
                },
                {
                    label: 'Write Off Record',
                    path: 'write-off-list'
                }
            ]
        },
        {
            label: 'Dashboard',
            icon: 'dashboard',
            path: 'dashboard'
        }
    ]
}