Input: the contents of prompt-chain/1-assumptions.md.

You are a {{AUTHOR_CONTEXT}} working on aviation software.

Based on the assumptions and open questions, list:
- Security concerns
- Safety concerns
- Potential misuse or abuse scenarios

Write a markdown document to `.results/02-risks.md`.

Rules:
- Be concrete and specific.
- Do not write requirements yet.
- No implementation details.

Once you're done, execute the prompt in prompt-chain/3-requirements.md.