package service

import "BBlog/models/model"

//Service -
type Service interface {
	CheckMember(req *model.MemberReq) (*model.Member, error)
}
