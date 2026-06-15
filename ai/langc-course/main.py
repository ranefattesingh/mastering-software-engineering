from dotenv import load_dotenv

load_dotenv()

from importlib.metadata import version
from langchain_ollama import ChatOllama

core_version = version("langchain_core")
lg_version = version("langgraph")

print(f"langchain-core version: {core_version}")
print(f"langgraph version: {lg_version}")


def main():
    #Test
    llm_gemma4 = ChatOllama(model="gemma4", temperature=0)
    response = llm_gemma4.invoke("Say 'setup complete!' in one word")
    print(f"Response from  ChatOllama: {response}")

    llm_llama3_2 = ChatOllama(model="llama3.2", temperature=0)
    response = llm_llama3_2.invoke("Say 'setup complete!' in one word")
    print(f"Response from  ChatOllama: {response}")

    print("Setup Complete")


if __name__ == "__main__":
    main()
