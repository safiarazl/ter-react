import React, {Component} from "react";
import Form from "../components/FormLogin";



export default class Login extends Component {
    constructor(props) {
      super(props)
    
      this.state = {
        username: "",
        password: "", 
        mergeusepass: []
      }
    }

    handleChange = (event) => {
        this.setState({
            [event.target.name]: event.target.value
        });
    }
    
    handleSubmit = (event) => {
        event.preventDefault();
        console.log("Data: ",this.state);
        this.setState({
            mergeusepass: [...this.state.mergeusepass, {
                username: this.state.username,
                password: this.state.password
            }]
        });
        this.setState ({
            username: "",
            password: ""
        });
    }

    render(){
        console.log(this.state.mergeusepass)
        return (
        <Form {...this.state} handleChange={this.handleChange} handleSubmit={this.handleSubmit}/>
    );}
}

// export default Login;