"""
Question 2
Challenge:
Set Up a Multi-Agent System with Manager and Web Search Agents

Assessment Criteria:
Web agent has correct tools configured
Manager agent properly references web agent
Appropriate max_steps value is set
Required imports are authorized
"""

# Create web agent and manager agent structure
web_agent = ToolCallingAgent(
    tools=[],           # Add required tools
    model=None,         # Add model
    max_steps=5,        # Adjust steps
    name="",           # Add name
    description=""      # Add description
)

manager_agent = CodeAgent()