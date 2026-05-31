"""
Question 1
Challenge:
Create a Basic Code Agent with Web Search Capability

Assessment Criteria:
Correct imports are included
DuckDuckGoSearchTool is added to tools list
HfApiModel is properly configured
Model ID is correctly specified
"""

# Create a CodeAgent with DuckDuckGo search capability
from smolagents import CodeAgent

agent = CodeAgent(
    tools=[],           # Add search tool here
    model=None          # Add model here
)