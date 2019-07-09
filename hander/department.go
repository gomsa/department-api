package hander

import (
	"context"

	"github.com/gomsa/department/client"
	departmentPB "github.com/gomsa/department/proto/department"
	"github.com/gomsa/tools/uitl"

	pb "github.com/gomsa/department-api/proto/department"
)

// Department 部门结构
type Department struct {
}

// List 部门列表
func (srv *Department) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &departmentPB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	departmentRes, err := client.Department.List(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(departmentRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取部门
func (srv *Department) Get(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	department := &departmentPB.Department{}
	err = uitl.Data2Data(req, department)
	if err != nil {
		return err
	}
	departmentRes, err := client.Department.Get(ctx, department)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(departmentRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建部门
func (srv *Department) Create(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	department := &departmentPB.Department{}
	err = uitl.Data2Data(req, department)
	if err != nil {
		return err
	}
	departmentRes, err := client.Department.Create(ctx, department)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(departmentRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新部门
func (srv *Department) Update(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	department := &departmentPB.Department{}
	err = uitl.Data2Data(req, department)
	if err != nil {
		return err
	}
	departmentRes, err := client.Department.Update(ctx, department)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(departmentRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除部门
func (srv *Department) Delete(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	department := &departmentPB.Department{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	departmentRes, err := client.Department.Delete(ctx, department)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(departmentRes, res)
	if err != nil {
		return err
	}
	return err
}
