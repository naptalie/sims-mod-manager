from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver.common.by import By
import time

"""
selenium webdriver class, can run as headless or not headless, can also stay_alive

PARAMS:
    headless: bool -> set whether to render browser
    stay_alive: bool -> keep the browser alive for debugging
    download_location: str -> set download location from agent
"""
class SeleniumAgent():

    def __init__(self, headless: bool = True, stay_alive: bool = True, download_location: str = "C:\\Users\\nick\\Downloads"):
        self.headless = headless
        self.stay_alive = stay_alive
        self.options = Options()
        self.download_location = download_location
        self.service = Service(ChromeDriverManager().install())
        self.__set_options()
        self.driver = webdriver.Chrome(service=self.service, options=self.options)


    """
    set the options for the webdriver

    PARAMS:
        None
    """
    def __set_options(self):
        prefs = {"download.default_directory" : f"{self.download_location}", 'directory_upgrade': True}
        if(self.headless):
            self.options.add_argument("--headless=chrome")
        if(self.stay_alive):
            self.options.add_experimental_option("detach", True)
        self.options.add_experimental_option("prefs",prefs)


    """
    invoke sleep

    PARAMS:
        sleep_time: int -> time to sleep
    """
    def agent_sleep(self, sleep_time: int) -> None:
        time.sleep(sleep_time)


    """
    supply a url for the webdriver to load

    PARAMS:
        returns: None
    """
    def get(self, url: str):
        self.driver.get(url)


    """
    supply a single html element and 'by' method and find the element

    PARAMS:
        by: By -> selenium action by which to locate element
        element: str -> corresponding element to look for
        returns: WebElement
    """
    def find_element(self, by: By, element: str) -> WebElement:
        return self.driver.find_element(by, element)

        
    """
    supply a single html element and 'by' method and find the element, optional sleep time

    PARAMS:
        by: By -> selenium action by which to locate element
        element: str -> corresponding element to look for
        sleep_before: int -> delay to time.sleep() before clicking
        returns: WebElement
    """
    def find_element_sleep(self, by: By, element: str, sleep_before: int = 0) -> WebElement:
        self.agent_sleep(sleep_before)
        return self.driver.find_element(by, element)