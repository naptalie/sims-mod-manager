from classes.selenium_agent import SeleniumAgent
from selenium.webdriver.common.by import By

agent = SeleniumAgent(headless=False, stay_alive=True)
url = "https://turbodriver.itch.io/wickedwhims"
agent.get(url)
agent.find_element_sleep(by=By.NAME, element='remember', sleep_before=2).click()
agent.find_element(by=By.XPATH, element='/html/body/div/div[1]/div/form/div[3]/button').click()
date = agent.find_element_sleep(by=By.CLASS_NAME, element='info_column', sleep_before=5).text
print(date)
agent.find_element(by=By.XPATH, element='/html/body/div/div/div[2]/div[2]/div[1]/div[3]/div/div[1]/a').click()
agent.agent_sleep(15)
agent.driver.close()