//
// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package annotations

const (
	vcAnnotationsPrefix = "com.github.containers.virtcontainers."

	// KernelPath is a pod annotation for passing a per container path pointing at the kernel needed to boot the container VM.
	KernelPath = vcAnnotationsPrefix + "KernelPath"

	// ImagePath is an pod annotation for passing a per container path pointing at the guest image that will run in the container VM.
	ImagePath = vcAnnotationsPrefix + "ImagePath"

	// KernelHash is a pod annotation for passing a container kernel image SHA-512 hash value.
	KernelHash = vcAnnotationsPrefix + "KernelHash"

	// ImageHash is an pod annotation for passing a container guest image SHA-512 hash value.
	ImageHash = vcAnnotationsPrefix + "ImageHash"

	// AssetHashType is the hash type used for assets verification
	AssetHashType = vcAnnotationsPrefix + "AssetHashType"
)

const (
	// SHA512 is the SHA-512 (64) hash algorithm
	SHA512 string = "sha512"
)
