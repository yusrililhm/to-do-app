import '../assets/css/index.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckSquare } from '@fortawesome/free-solid-svg-icons';
import { faXmarkSquare } from '@fortawesome/free-solid-svg-icons';
import { faPenSquare } from '@fortawesome/free-solid-svg-icons';
import { Component } from 'react';
import axios from 'axios';

let endpoint = "http://localhost:8000";

class ToDoList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            todo: "",
            items: []
        }
    };
    componentDidMount() {
        this.showTodo();
    }

    onChange = event => {
        this.setState({
            [event.target.name]: event.target.value
        })
    }

    onSubmit = () => {
        let { todo } = this.state;
        if (todo) {
            axios
            .post(
                endpoint + "/api/insert",
                {
                    todo
                },
                {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    }
                }
            )
            .then (res => {
                this.showTodo();
                this.setState({
                    todo: ""
                });
                console.log(res)
            })
        }
    }
    showTodo = () => {
        axios.get(endpoint + "/a[i/show").then(res => {
            console.log(res.data)
            if (res.data) {
                this.setState({
                    items: res.data.map(item => {
                        return (
                            <section className='font-Poppins h-screen'>
                            <table className='m-auto mt-4 justify-center items-center shadow-lg bg-white w-3/5 h-20'>
                                {/* <thead>
                                    <tr>
                                        <td className='text-center'>Todo</td>
                                        <td className='text-center'>Actions</td>
                                    </tr>
                                </thead> */}
                                <tbody className=''>
                                <tr key={item._id}>
                                    <td className='w-5/6'>{item.todo}</td>
                                    <td className='text-center'>
                                        <a href='/done' className='text-green-500 mx-1'><FontAwesomeIcon icon={faCheckSquare} /></a>
                                        <a href='/update' className='text-yellow-500 mx-1'><FontAwesomeIcon icon={faPenSquare} /></a>
                                        <a href='/delete' className='text-red-500 mx-1'><FontAwesomeIcon icon={faXmarkSquare} /></a>
                                    </td>
                                </tr>
                                </tbody>
                            </table>
                        </section>
                        )
                    })
                })
            } else {
                this.setState({
                    items: []
                })
            }
        })
    };
    render() {
        return (
            <div className='flex justify-center items-center'>
                <form onSubmit={this.onSubmit} action='/api/add' className='flex justify-center items-center mt-20 bg-white shadow-lg py-2 rounded-md w-3/5'>
                    <input type='text' name='todo' id='todo' placeholder='Type your plan here' className='mx-4 w-full' onChange={this.onChange} value={this.state.todo}/>
                </form>
            </div>
        )
    }
};

export default ToDoList;