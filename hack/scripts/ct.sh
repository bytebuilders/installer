#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eou pipefail

for dir in charts/*/; do
    dir=${dir%*/}
    dir=${dir##*/}
    num_files=$(find charts/${dir}/templates -type f | wc -l)
    echo $dir
    if [ $num_files -le 1 ] || [[ "$dir" = "accounts-ui" ]] || [[ "$dir" = "ace" ]] || [[ "$dir" = "billing" ]] || [[ "$dir" = "grafana" ]] || [[ "$dir" = "platform-api" ]] || [[ "$dir" = "prom-proxy" ]] || [[ "$dir" = "opscenter-features" ]] || [[ "$dir" = "dns-proxy" ]] || [[ "$dir" = "smtprelay" ]]; then
        make ct CT_COMMAND=lint TEST_CHARTS=charts/$dir
    else
        make ct TEST_CHARTS=charts/$dir
    fi
done
