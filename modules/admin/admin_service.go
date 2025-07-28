package admin

// Admin Interface
type AdminService interface {
}

// Admin Struct
type adminService struct {
	adminRepo AdminRepository
}

// Admin Constructor
func NewAdminService(adminRepo AdminRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}
