from selenium.webdriver.chrome.options import Options    
 
    options = Options()
    options.add_argument("--headless")
    options.add_argument("--start-maximized")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-extensions")
    options.add_argument('--disable-dev-shm-usage')    
    options.add_argument("--disable-gpu")
    options.add_argument('--disable-software-rasterizer')
    options.add_argument("user-agent=Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 640 XL LTE) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.10166")
    options.add_argument("--disable-notifications")

    options.add_experimental_option("prefs", {
        "download.default_directory": "/home/test",
        "download.prompt_for_download": False,
        "download.directory_upgrade": True,
        "safebrowsing_for_trusted_sources_enabled": False,
        "safebrowsing.enabled": False
        }
    )