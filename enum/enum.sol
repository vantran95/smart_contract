pragma solidity >=0.7.0 < 0.9.0;

contract Enum {
    enum ActionChoices {GoLeft, GoRight, GoStraight, SitStill}

    ActionChoices choice;
    ActionChoices constant defaultChoice = ActionChoices.GoStraight;

    function setGoStraight() {
        choice = ActionChoices.GoStraight;
    }

    function getChoice() returns (ActionChoices) {
        return choice;
    }

    function getDefaultChoice() returns (uint) {
        return uint(defaultChoice);
    }
}