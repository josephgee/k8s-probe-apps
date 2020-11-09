# k8s-probe-apps

This projects exist to reproduce certain application behavior in a kubernetes environment. Each has a `build.sh` script
that will:
- build a probe app
- containerize it
- push to docker hub
- apply a pod to run the app in the current cluster

## crash-loop

This app crashes on a regular interval to observe termination status behavior

## cpu-loader

This app ramps up cpu load gradually to 100% and back down again to observe cpu load reporting
