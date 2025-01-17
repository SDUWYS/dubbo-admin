// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identifier

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of dubbo-admin
	// This relies on the fact this file is 3 levels up from the root; if this changes, adjust the path below
	Root            = filepath.Join(filepath.Dir(b), "../../..")
	Deploy          = filepath.Join(Root, "/deploy")
	Charts          = filepath.Join(Deploy, "/charts")
	Profiles        = filepath.Join(Deploy, "profiles")
	Addons          = filepath.Join(Deploy, "addons")
	AddonDashboards = filepath.Join(Addons, "dashboards")
	AddonManifests  = filepath.Join(Addons, "manifests")
)
