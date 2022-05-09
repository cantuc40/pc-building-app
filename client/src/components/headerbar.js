import {Component} from 'react';
import {Link} from 'react-router-dom';




export default class HeaderBar extends Component {

    render(){
        return (
        <nav>
            <div>
                <p>PC Builder</p>
                <button type="button">
                    <span></span>
                </button>
                <div>
                    <ul>
                        <li>
                            <Link to='/' className="nav-link">Home</Link>
                        </li>
                        <li>
                            <Link to='/pcs'>PCs</Link>
                        </li>
                        <li>
                            <Link to='/parts'>Parts</Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        )
    }
}