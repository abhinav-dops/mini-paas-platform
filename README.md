## App Contract

Every app deployed to Mini PaaS must:

1. Contain a Dockerfile
2. Expose one HTTP port (e.g., 8080)
3. Be stateless (no local persistence)
4. Start with one container
5. Bind to 0.0.0.0
