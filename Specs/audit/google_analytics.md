# Overview

This spec describes the list of Google Analytics Audits we would conduct when auditing Google Analytics

- Profile
    - Is there more than 1 profile for the account?
    - Is there a profile that has no filters attached to it? (a Raw profile)
- Custom Dimensions
    - Are there custom dimensions being set?
    - Are the custom dimensiosn being used? 
        - Indicated by data collected based on it
- Custom Metrics
    - Are there custom metrics being set?
    - Are the custom metrics being used? 
        - Indicated by data collected based on it
- Goals
    - Are there goals set
    - Are the goals being used? 
        - Indicated by data collected based on it
    - Are there duplicate goals being set?
        - Check every node in the goal settings
- Are display advertisement reports being used? 
    - Allow display of Age/Gender demographic reports
        - Check the age, gender, demographic reports if there is any data there
- Traffic Sources
    - Is the source, medium, campaign inconsistent?
        - Indicated by collecting data and checking that all are small caps