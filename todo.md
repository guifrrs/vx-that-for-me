# Todo - Code Improvements

## Critical Bug Fixes
1. [x] Fix unhandled errors in handlers.go (SendMessage and DeleteMessage on lines 38, 44)
2. [x] Fix URL handling bug - support http:// (not just https://)
3. [x] Add TELEGRAM_TOKEN validation before starting bot
4. [x] Fix MessageHandler regex mismatch with replaceLink logic

## Performance Improvements
5. [x] Pre-compile regex patterns as package-level variables (avoid recompilation on every call)

## Missing Tests
6. [x] Add tests for getUsername() function
7. [x] Add more edge case tests for replaceLink()
   - URLs with www. subdomain
   - URLs with query parameters
   - URLs with anchors
   - http:// protocol
   - twitter.com domain
8. [-] Add integration tests for MessageHandler() - Skipped (requires mocking tbot.Client)

## Code Quality
9. [x] Extract fixupx.com to a constant
10. [x] Improve logging with structured format (log.Printf)
11. [-] Refactor global state into encapsulated structure - Partial (kept simple for now)
12. [-] Use net/url for URL parsing instead of regex where appropriate - Regex is sufficient for this use case

## Security
13. [x] Sanitize logged user messages (strip control chars and dangerous content)
14. [-] Add basic rate limiting or message length validation - Out of scope for simple bot

## Dockerfile Improvements
15. [x] Update alpine version from 3.16 to 3.20
16. [x] Add non-root user for security
17. [x] Add health check

## Infrastructure
18. [x] Add graceful shutdown handling
19. [-] Add configuration validation struct - Not needed for single env var
20. [x] Verify go.sum is tracked in git

## Summary

All critical bugs and high-priority improvements have been completed:
- Fixed 4 critical bugs
- Added 15 new test cases
- Improved security with input sanitization
- Modernized Dockerfile with non-root user and health checks
- Added graceful shutdown with signal handling
- go.sum verified in git

Remaining low-priority items:
- Integration tests for MessageHandler() - would require mocking framework
- Rate limiting - not critical for simple bot
- Configuration struct - overkill for single env var
