# emilynardoni.com

The deployment process is pretty low tech. After checking out the repository, assuming the fly.io cli is installed, run `fly launch` to go through the entire build + deploy process interactively. By doing so, one can re-configure the deployment settings.

Alternatively, to simply update the deployment using the existing deployment settings (for example after making some small update to the application), simply run `fly deploy`.

Note: If you update CSS, bump the cache-buster query parameter in `static/html/base.layout.tmpl` (the `?v=...` on the stylesheet link) so browsers fetch the new version.
