import json
import urllib.request
import urllib.error

def get_version_data(url):
    try:
        with urllib.request.urlopen(url) as response:
            data = json.loads(response.read().decode())
            return data.get("current_version")
    except urllib.error.URLError as e:
        print(f"Failed to fetch data from URL: {url}. Error: {e}")
        return None

def read_local_version_file(file_path):
    try:
        with open(file_path, 'r') as file:
            data = json.load(file)
            return data.get("current_version")
    except FileNotFoundError as e:
        print(f"Failed to open file: {file_path}. Error: {e}")
        return None

def main():
    
    file_path = "current_version.json"
    version_from_file = read_local_version_file(file_path)
    
    url = "https://theyonk.gateway.scarf.sh/callhome/version.json/" + version_from_file
    version_from_url = get_version_data(url)
    

    if version_from_url is None or version_from_file is None:
        print("Could not fetch version information.")
        return

    if version_from_url == version_from_file:
        print("Version match. The version is: ", version_from_url)
    else:
        print("Version mismatch. URL version: ", version_from_url, ". File version: ", version_from_file)

if __name__ == "__main__":
    main()
