"""
Question 3
Challenge:
Configure Agent Security Settings

Assessment Criteria:
E2B sandbox is properly configured
Authorized imports are appropriately limited
Security settings are correctly implemented
Basic agent configuration is maintained
"""
# Set up secure code execution environment
from smolagents import CodeAgent

agent = CodeAgent(
    tools=[],
    model=model
    # Add security configuration
)