import {Component} from 'react';
import { useQuery, gql } from "@apollo/client";





export default class Parts extends Component {
    constructor(props) {
        super(props);
        this.state = {
            parts: [],
            partName: this.props.partName
        }
    }

    componentDidMount(){

        const PARTS_QUERY = gql
        `{
            ${this.state.partName}{

            }
        }`;

        const {data, loading, err} = useQuery(PARTS_QUERY);


    }
    


    render(){
        return (
            <div>
                
            </div>
        )
    }
}