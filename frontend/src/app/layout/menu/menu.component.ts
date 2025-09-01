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
      /*  this.userService.menu$.subscribe(menuItems => {
            this.menuItems = menuItems
        }) */
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
            name: 'Home',
            icon: 'home'
        },
        {
            name: 'System Management',
            icon: 'settings',
            isOpen: false,
            childrens: [
                {
                    name: 'User',
                    path: 'users',
                    icon: 'account-circle'
                },
                {
                    name: 'Role',
                    path: 'role',
                    icon: ''
                },
                {
                    name: 'Menu',
                    icon: '',
                    path: 'menu'
                },
                {
                    name: 'Department',
                    path: 'departments',
                    icon: 'team_dashboard'
                }
            ]
        },
        {
            name: 'Base Management',
            icon: 'dock_to_bottom',
            isOpen: false,
            childrens: [
                {
                    name: 'Tax Information',
                    icon: '',
                    path: 'tax-information'
                },
                {
                    name: 'Budget',
                    path: 'budget',
                    icon: ''
                },
                {
                    name: 'Action Log',
                    icon: '',
                    path: 'action-record'
                },
                {
                    name: 'Vendor',
                    icon: '',
                    path: 'vendor'
                },
                {
                    name: 'Location',
                    path: 'location',
                    icon: 'location'
                },
                {
                    name: 'Code Type',
                    icon: '',
                    path: 'code-type'
                }
            ]
        },
        {
            name: 'Asset Managemnt',
            icon: 'save',
            childrens: [
                {
                    name: 'Asset Type',
                    icon: '',
                    path: 'asset-type'
                },
                {
                    name: 'Asset List',
                    icon: '',
                    path: 'asset-lists'
                },
                {
                    name: 'Create Asset',
                    icon: '',
                    path: 'asset-create'
                },
                {
                    name: 'Stock Take',
                    icon: '',
                    path: 'stock-takes'
                },
                {
                    name: 'Asset List Report',
                    icon: '',
                    path: 'asset-list-all'
                },
                {
                    name: 'Inventory Record',
                    icon: '',
                    path: 'inventory-record'
                },
                {
                    name: 'Repair Record',
                    icon: '',
                    path: 'repair-records'
                },
                {
                    name: 'Write Off Record',
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