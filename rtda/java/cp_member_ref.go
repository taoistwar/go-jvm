package java

import "github.com/taoistwar/go-jvm/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (mr *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	mr.className = refInfo.ClassName()
	mr.name, mr.descriptor = refInfo.NameAndDescriptor()
}

func (mr *MemberRef) Name() string {
	return mr.name
}
func (mr *MemberRef) Descriptor() string {
	return mr.descriptor
}
