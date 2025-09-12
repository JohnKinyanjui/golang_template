- All APIs must be compiled into a single Insomnia collection at `./docs/insomnia.yaml`.
- Always reuse the same `_id` values for workspace, requests, and environments to avoid duplicates; re-importing will then update the collection instead of creating a new one.
- The collection must define `http://localhost:8000` as the base URL and use an environment variable `{{TOKEN}}` for the `Authorization: Bearer {{TOKEN}}` header.

- To update the workspace, run:
  ```bash
  inso import spec ./docs/insomnia.yaml --force
  ```
