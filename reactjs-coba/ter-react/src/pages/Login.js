import React, { Component } from "react";
import { Await } from "react-router-dom";
import Form from "../components/FormLogin";

export default class Login extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: "",
      password: "",
      mergeusepass: [],
    };
  }

  handleChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  handleSubmit = (event) => {
    event.preventDefault();
    // console.log("Data: ", this.state);
    console.log("username: ", this.state.username);
    console.log("password: ", this.state.password);
    // new
    let data = {
      username: this.state.username,
      password: this.state.password,
    };
    const response = async () => {
      fetch("http://localhost:8080/user/login", {
        method: "POST",
        mode: "no-cors",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      var myJson = await response.json();
      return myJson;
    };
    // end new
    this.setState({
      mergeusepass: [
        ...this.state.mergeusepass,
        {
          username: this.state.username,
          password: this.state.password,
        },
      ],
    });
    this.setState({
      username: "",
      password: "",
    });
  };

  render() {
    console.log(this.state.mergeusepass);
    return (
      <Form
        {...this.state}
        handleChange={this.handleChange}
        handleSubmit={this.handleSubmit}
      />
    );
  }
}

// export default Login;
