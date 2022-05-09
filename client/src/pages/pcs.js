import {Component} from 'react';
import { Link } from 'react-router-dom'; //Link to Each PC

//GraphQL Query for retrieving PCs



export default class PCS extends Component {
    constructor(props){
        super(props);
        this.state= {
            PCs: []
        }
    }

    componentDidMount(){}

    render(){
        return (
            <div>
                
            </div>
        )
    }
}