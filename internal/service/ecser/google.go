package ecser

//
// import (
// 	"context"
// 	"fmt"
//
// 	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
// 	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
// 	"github.com/pkg/errors"
// 	"google.golang.org/api/compute/v1"
// 	"google.golang.org/api/option"
// )
//
// // Google当前存在2个问题：
// // 1. 权限的最佳实践
// // 2. 国内无法直接访问google认证的接口
// type GoogleEcs struct {
// 	cli *compute.Service
// }
//
// func NewGoogleEcsClient(tenant tenanter.Tenanter) (Ecser, error) {
// 	var client *compute.Service
//
// 	// rName, err := tenanter.GetAwsRegionName(regionId)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
//
// 	client, err := compute.NewService(context.Background(), option.WithCredentialsFile("/Users/didi/Study/cloud-fitter/google_auth.json"))
// 	if err != nil {
// 		return nil, errors.WithMessage(err, "new compute service error")
// 	}
//
// 	// // Project ID for this request.
// 	// project := "my-project" // TODO: Update placeholder value.
// 	//
// 	// // The name of the zone for this request.
// 	// zone := "my-zone" // TODO: Update placeholder value.
// 	//
// 	// req := computeService.Instances.ListDetail(project, zone)
// 	// if err := req.Pages(ctx, func(page *compute.InstanceList) error {
// 	// 	for _, instance := range page.Items {
// 	// 		// TODO: Change code below to process each `instance` resource:
// 	// 		fmt.Printf("%#v\n", instance)
// 	// 	}
// 	// 	return nil
// 	// }); err != nil {
// 	// 	log.Fatal(err)
// 	// }
//
// 	// switch t := tenant.(type) {
// 	// case *tenanter.AccessKeyTenant:
// 	// 	auth := basic.NewCredentialsBuilder().WithAk(t.GetId()).WithSk(t.GetSecret()).Build()
// 	// 	hcClient := hwecs.EcsClientBuilder().WithRegion(region.ValueOf(rName)).WithCredential(auth).Build()
// 	// 	client = hwecs.NewEcsClient(hcClient)
// 	// default:
// 	// }
//
// 	if err != nil {
// 		return nil, errors.Wrap(err, "init google ecs client error")
// 	}
// 	return &GoogleEcs{cli: client}, nil
// }
//
// // func (ecs *GoogleEcs) ECSStatistic() (*pbecs.ECSStatisticRespList, error) {
// // 	return nil, nil
// // }
//
// func (ecs *GoogleEcs) ListDetail(pageNumber, pageSize int) (*pbecs.ListResp, error) {
// 	// req := new(model.ListServersDetailsRequest)
// 	// offset := int32((pageNumber - 1) * pageSize)
// 	// req.Offset = &offset
// 	// limit := int32(pageSize)
// 	// req.Limit = &limit
// 	//
// 	// resp, err := ecs.cli.ListServersDetails(req)
// 	// if err != nil {
// 	// 	return nil, errors.Wrap(err, "Google ListDetail error")
// 	// }
//
// 	// Project ID for this request.
// 	project := "focused-stacker-311609" // TODO: Update placeholder value.
//
// 	// The name of the zone for this request.
// 	zone := "asia-east2-a" // TODO: Update placeholder value.
//
// 	req := ecs.cli.Instances.ListDetail(project, zone)
// 	var ecses []*pbecs.EcsInstance
// 	if err := req.Pages(context.Background(), func(page *compute.InstanceList) error {
// 		for _, instance := range page.Items {
// 			ecses = append(ecses, &pbecs.EcsInstance{
// 				InstanceId:   fmt.Sprint(instance.Id),
// 				InstanceName: instance.AccountName,
// 				RegionId:     "",
// 				ZoneId:       "",
// 				PublicIps:    nil,
// 				Status:       "",
// 				CreationTime: "",
// 				Description:  "",
// 			})
// 		}
// 		return nil
// 	}); err != nil {
// 		return nil, err
// 	}
//
// 	return &pbecs.ListResp{
// 		Ecses: ecses,
// 	}, nil
// }
