import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "https://localhost:9000"

class ToDoList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            task: "",
            items: []
        };
    }

    ComponentDidMount() {
        this.getTask();
    }


    render() {
        return (
            <div className="row">
                <div>
                    <Header className="header" as="h2" color="yellow">
                        TO DO LIST
                    </Header>
                </div>
            </div>
        );
    }
}



export default ToDoList;