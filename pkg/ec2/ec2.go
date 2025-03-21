package ec2

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func CreateInstance(cfg aws.Config, ami, instanceType, key, subnet, instanceProfile string, count, diskSize int, sg []string, tags map[string]string) {
	client := ec2.NewFromConfig(cfg)
	ctx := context.TODO()

	var tag []types.Tag
	var tagArray []types.TagSpecification

	t := types.TagSpecification{
		ResourceType: "instance",
		Tags:         tag,
	}

	tagArray = append(tagArray, t)

	profile := types.IamInstanceProfileSpecification{
		Name: &instanceProfile,
	}

	input := &ec2.RunInstancesInput{
		InstanceType:       types.InstanceType(instanceType),
		ImageId:            &ami,
		KeyName:            &key,
		SecurityGroupIds:   sg,
		IamInstanceProfile: &profile,
		SubnetId:           &subnet,
		TagSpecifications:  tagArray,
	}

	i, err := client.RunInstances(ctx, input)

	// TODO: figure out which SDK error messages can be used here.
	if err != nil {
		log.Fatal("Failed to run EC2 instance.")
	}

	var instanceId []string
	for _, instance := range i.Instances {
		instanceId = append(instanceId, *instance.InstanceId)
	}

	for id := range instanceId {
		fmt.Println("Launched instance with Id:", id)
	}
}

func DeleteInstance(cfg aws.Config, i string) {
	client := ec2.NewFromConfig(cfg)
	ctx := context.TODO()

	var ids []string
	ids = append(ids, i)

	input := &ec2.TerminateInstancesInput{
		InstanceIds: ids,
	}

	out, err := client.TerminateInstances(ctx, input)

	if err != nil {
		fmt.Println(err)
	}

	for _, t := range out.TerminatingInstances {
		var terminatedCode *int32
		status := t.CurrentState.Code

		for {
			time.Sleep(10 * time.Second)
			if status == terminatedCode {
				fmt.Println("Instance with Id" + i + "successfully terminated")
				break
			}
		}
	}
}
