Input:

- prompt-chain/1-assumptions.md
- prompt-chain/2-security.md

You are a {{AUTHOR_CONTEXT}}.

Generate a short requirements document for a {{FEATURE_NAME}} in output.md that will be implemented by a {{TARGET_ENGINEER}}. Emphasize SECURITY.

Before writing, create a short checklist of what the document must include, and use it to ensure nothing is omitted.

Write the document in markdown with the following structure:
- Checklist
- Assumptions
- Overview
- Functional Requirements
- Non-Functional Requirements
- Open Questions

Rules:
- Each requirement must have an ID (e.g., TODO-1, TODO-2).
- Requirements must be testable and measurable.
- Do not include implementation code or specific libraries.
- Ensure requirements address the security and misuse concerns from prompt-chain/2-security.md.