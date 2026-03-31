import os
import json
import logging
from typing import Dict, List, Optional

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def load_json_config(file_path: str) -> Dict[str, Optional[Dict[str, str]]]:
    """
    Load JSON configuration from file.
    
    Args:
    file_path: Path to the JSON file.
    
    Returns:
    Dictionary containing the configuration.
    """
    try:
        with open(file_path, 'r') as file:
            return json.load(file)
    except FileNotFoundError:
        logger.error(f"Config file {file_path} not found.")
        return {}
    except json.JSONDecodeError as e:
        logger.error(f"Failed to parse config file {file_path}: {e}")
        return {}

def get_env_var(var_name: str) -> Optional[str]:
    """
    Get environment variable.
    
    Args:
    var_name: Name of the environment variable.
    
    Returns:
    Value of the environment variable, or None if it's not set.
    """
    return os.getenv(var_name)

def list_dir(path: str) -> List[str]:
    """
    List files and directories in a directory.
    
    Args:
    path: Path to the directory.
    
    Returns:
    List of files and directories in the directory.
    """
    return os.listdir(path)

def create_dir(path: str) -> None:
    """
    Create a directory.
    
    Args:
    path: Path to the directory.
    
    Returns:
    None
    """
    os.makedirs(path, exist_ok=True)