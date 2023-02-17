import '../assets/css/index.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckSquare } from '@fortawesome/free-solid-svg-icons';
import { faXmarkSquare } from '@fortawesome/free-solid-svg-icons';
import { faPenSquare } from '@fortawesome/free-solid-svg-icons';
import { Component } from 'react';
import axios from 'axios';
import { Card } from 'semantic-ui-react';

let endpoint = "http://localhost:8000";

class ToDo extends Component {
    constructor(props) {
        super(props);

        this.state = {
            Todo: "",
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
        let { Todo } = this.state;
        if (Todo) {
            axios
            .post(
                endpoint + "/api/insert",
                {
                    Todo
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
                    Todo: ""
                });
                console.log(res)
            })
        }
    }
    showTodo = () => {
        axios.get(endpoint + "/api/show").then(res => {
            console.log(res.data)
            if (res.data) {
                this.setState({
                    items: res.data.map(item => {
                        return (
                            <Card key={item.ID}>
                                <Card.Content className='flex justify-center items-center w-3/5 bg-white my-4 text-black h-16 shadow-md rounded-md m-auto'>
                                    <Card.Header className='w-5/6'>
                                        <div className='mx-3'>{item.Todo}</div>
                                    </Card.Header>
                                    <Card.Meta>
                                        <a href='/done' className='text-green-500 mx-1'><FontAwesomeIcon icon={faCheckSquare} /></a>
                                        <a href='/update' className='text-yellow-500 mx-1'><FontAwesomeIcon icon={faPenSquare} /></a>
                                        <a href='/delete' className='text-red-500 mx-1'><FontAwesomeIcon icon={faXmarkSquare} /></a>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                            // <section className='font-Poppins h-screen'>
                            //     <table className='m-auto mt-4 justify-center items-center shadow-lg bg-white w-3/5 h-20' key={item._id}>
                            //         <tbody className='' key={item._id}>
                            //         <tr key={item._id}>
                            //             <td className='w-5/6'>{item.todo}</td>
                            //             <td className='text-center'>
                            //                 <a href='/done' className='text-green-500 mx-1'><FontAwesomeIcon icon={faCheckSquare} /></a>
                            //                 <a href='/update' className='text-yellow-500 mx-1'><FontAwesomeIcon icon={faPenSquare} /></a>
                            //                 <a href='/delete' className='text-red-500 mx-1'><FontAwesomeIcon icon={faXmarkSquare} /></a>
                            //             </td>
                            //         </tr>
                            //         </tbody>
                            //     </table>
                            // </section>
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
            <section className='h-screen'>
                <div className='flex justify-center items-center'>
                    <form onSubmit={this.onSubmit} className='flex justify-center items-center mt-20 bg-white shadow-lg py-2 rounded-md w-3/5'>
                        <input 
                        type='text' 
                        name='Todo' 
                        id='todo' 
                        placeholder='Type your plan here' 
                        className='w-full mx-4' 
                        onChange={this.onChange}
                        />
                    </form>
                </div>
                <div className=''>
                    <Card.Group className=''>{this.state.items}</Card.Group>
                </div>
            </section>
        )
    }
};

export default ToDo;