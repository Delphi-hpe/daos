//
// (C) Copyright 2018-2020 Intel Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
// The Government's rights to use, modify, reproduce, release, perform, display,
// or disclose this software are subject to the terms of the Apache License as
// provided in Contract No. 8F-30005.
// Any reproduction of computer software, computer software documentation, or
// portions thereof marked with this legend must also reproduce the markings.
//
// Package code is a central repository for all control plane fault codes.
package code

// Code represents a stable fault code.
//
// NB: All control plane errors should register their codes in the
// following block in order to avoid conflicts.
//
// Also note that new codes should always be added at the bottom of
// their respective blocks. This ensures stability of fault codes
// over time.
type Code int

const (
	// general fault codes
	Unknown Code = iota
	MissingSoftwareDependency
	PrivilegedHelperNotPrivileged
	PrivilegedHelperNotAvailable
	PrivilegedHelperRequestFailed

	// generic storage fault codes
	StorageUnknown Code = iota + 100
	StorageAlreadyFormatted
	StorageFilesystemAlreadyMounted
	StorageDeviceAlreadyMounted
	StorageTargetAlreadyMounted

	// SCM fault codes
	ScmUnknown Code = iota + 200
	ScmFormatInvalidSize
	ScmFormatInvalidDeviceCount
	ScmFormatMissingMountpoint
	ScmFormatMissingDevice
	ScmFormatMissingParam
	ScmFormatConflictingParam
	ScmDiscoveryFailed

	// Bdev fault codes
	BdevUnknown Code = iota + 300
	BdevFormatUnknownClass
	BdevFormatFailure
	BdevFormatBadPciAddress

	// DAOS system fault codes
	SystemUnknown Code = iota + 400
	SystemMemberExists
	SystemMemberMissing
	SystemMemberChanged

	// client fault codes
	ClientUnknown Code = iota + 500
	ClientConfigBadControlPort
	ClientConfigBadAccessPoints

	// server fault codes
	ServerUnknown Code = iota + 600
	ServerBadConfig
	ServerNoConfigPath
	ServerConfigBadControlPort
	ServerConfigBadAccessPoints
	ServerConfigBadProvider
	ServerConfigNoServers
	ServerScmUnmanaged
	ServerBdevNotFound
	ServerConfigDuplicateFabric
	ServerConfigDuplicateLogFile
	ServerConfigDuplicateScmMount
	ServerConfigDuplicateScmDeviceList
	ServerConfigOverlappingBdevDeviceList
	ServerIommuDisabled
	ServerPoolScmTooSmall
	ServerPoolNvmeTooSmall

	// spdk library bindings codes
	SpdkUnknown Code = iota + 700
	SpdkCtrlrNoHealth
	SpdkBindingRetNull
	SpdkBindingFailed

	// security fault codes
	SecurityUnknown Code = iota + 900
)
