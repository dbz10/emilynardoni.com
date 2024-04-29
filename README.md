# emilynardoni.com

The deployment process is pretty low tech. After checking out the repository, assuming the fly.io cli is installed, run `fly launch` to go through the entire build + deploy process interactively. By doing so, one can re-configure the deployment settings.

Alternatively, to simply update the deployment using the existing deployment settings (for example after making some small update to the application), simply run `fly deploy`.
