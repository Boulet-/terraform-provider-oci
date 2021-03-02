// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_artifacts "github.com/oracle/oci-go-sdk/v36/artifacts"
)

func init() {
	RegisterResource("oci_artifacts_container_configuration", ArtifactsContainerConfigurationResource())
}

func ArtifactsContainerConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createArtifactsContainerConfiguration,
		Read:     readArtifactsContainerConfiguration,
		Update:   updateArtifactsContainerConfiguration,
		Delete:   deleteArtifactsContainerConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_repository_created_on_first_push": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createArtifactsContainerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).artifactsClient()

	return CreateResource(d, sync)
}

func readArtifactsContainerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).artifactsClient()

	return ReadResource(sync)
}

func updateArtifactsContainerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).artifactsClient()

	return UpdateResource(d, sync)
}

func deleteArtifactsContainerConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ArtifactsContainerConfigurationResourceCrud struct {
	BaseCrud
	Client                 *oci_artifacts.ArtifactsClient
	Res                    *oci_artifacts.ContainerConfiguration
	DisableNotFoundRetries bool
}

func (s *ArtifactsContainerConfigurationResourceCrud) ID() string {
	return s.D.Get("compartment_id").(string)
}

func (s *ArtifactsContainerConfigurationResourceCrud) Create() error {
	request := oci_artifacts.UpdateContainerConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isRepositoryCreatedOnFirstPush, ok := s.D.GetOkExists("is_repository_created_on_first_push"); ok {
		tmp := isRepositoryCreatedOnFirstPush.(bool)
		request.IsRepositoryCreatedOnFirstPush = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.UpdateContainerConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerConfiguration
	return nil
}

func (s *ArtifactsContainerConfigurationResourceCrud) Get() error {
	request := oci_artifacts.GetContainerConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.GetContainerConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerConfiguration
	return nil
}

func (s *ArtifactsContainerConfigurationResourceCrud) Update() error {
	request := oci_artifacts.UpdateContainerConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isRepositoryCreatedOnFirstPush, ok := s.D.GetOkExists("is_repository_created_on_first_push"); ok {
		tmp := isRepositoryCreatedOnFirstPush.(bool)
		request.IsRepositoryCreatedOnFirstPush = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.UpdateContainerConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerConfiguration
	return nil
}

func (s *ArtifactsContainerConfigurationResourceCrud) SetData() error {
	if s.Res.IsRepositoryCreatedOnFirstPush != nil {
		s.D.Set("is_repository_created_on_first_push", *s.Res.IsRepositoryCreatedOnFirstPush)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	return nil
}
