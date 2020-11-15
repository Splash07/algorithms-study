COMMAND Design Pattern

What it does:
  - Focus on encapsulation and abstraction
  - Not need to know the details and how the functions are implemented, instead have higher-level objects to command (encapsulate) those functions

In this example:
  - Command:
    - There are four commands one can use: On, Off, Increase Volume, Decrease Volume. 
    - Each command is associated with a device (which is TV in this case)


  - Device:
    - TV provides those commands with concrete implementations
  
  
  - Button:
    - Each button is associated with a command
    - Button objects encapsulate the commands and provide only one method (press) user can interact with
    - That way, button can use command (press) to invoke method from device without knowing the details

UML diagram:
![alt text](https://miro.medium.com/max/700/0*0ax_IClv9aVFe7i9.png "UML Diagram")


References:
  - https://levelup.gitconnected.com/the-command-pattern-with-go-fd5dabc84c7
