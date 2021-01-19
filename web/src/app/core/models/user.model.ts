import { AccountModel } from './account.model';
import { MembershipDetailsModel } from './membership-details.model';
import { RolePermissionsModel } from './role-permissions.model';

export class UserModel {
    ID?: number;
    FirstName?: string;
    LastName?: string;
    Email?: string; 
    CreatedAt?: string;
    DeletedAt?: string;
    ImageUrl?: string;
    LastLoggedInAt?: string;
    UpdatedAt?: string;
    Account?: AccountModel;
    MembershipDetails?: MembershipDetailsModel;
    RolePermissions?: RolePermissionsModel;
}
