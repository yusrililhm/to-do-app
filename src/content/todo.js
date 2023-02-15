import '../assets/css/index.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleCheck } from '@fortawesome/free-solid-svg-icons';
import { faCircleXmark } from '@fortawesome/free-solid-svg-icons';
import { faPen } from '@fortawesome/free-solid-svg-icons';

const ToDoList = () => {
    return (
        <section className='font-Poppins h-screen'>
            <div className='flex justify-center items-center'>
                <form method='POST' action='/Add' className='flex justify-center items-center mt-20 bg-white shadow-lg py-2 px-4 rounded-full w-3/5'>
                    <input type='text' name='todo' id='todo' placeholder='Type your plan here' className='w-full'/>
                </form>
            </div>
            <table>
                <thead>
                    <tr>
                        <td className='text-center'>Todo</td>
                        <td className='text-center'>Actions</td>
                    </tr>
                </thead>
                <tbody>
                <tr>
                    <td>{}</td>
                    <td>
                        <a href='/done' className='text-green-500 mx-1'><FontAwesomeIcon icon={faCircleCheck} /></a>
                        <a href='/update' className='text-yellow-500 mx-1'><FontAwesomeIcon icon={faPen} /></a>
                        <a href='/delete' className='text-red-500 mx-1'><FontAwesomeIcon icon={faCircleXmark} /></a>
                    </td>
                </tr>
                </tbody>
            </table>
        </section>
    );
}

export default ToDoList;