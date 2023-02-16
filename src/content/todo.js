import '../assets/css/index.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckSquare } from '@fortawesome/free-solid-svg-icons';
import { faXmarkSquare } from '@fortawesome/free-solid-svg-icons';
import { faPenSquare } from '@fortawesome/free-solid-svg-icons';

const ToDoList = () => {
    return (
        <section className='font-Poppins h-screen'>
            <div className='flex justify-center items-center'>
                <form method='POST' action='/add' className='flex justify-center items-center mt-20 bg-white shadow-lg py-2 rounded-md w-3/5'>
                    <input type='text' name='todo' id='todo' placeholder='Type your plan here' className='mx-4 w-full'/>
                    {/* <button className='text-center bg-teal-500 text-white font-bold'>Add Notes</button> */}
                </form>
            </div>
            <table className='m-auto mt-4 justify-center items-center shadow-lg bg-white w-3/5 h-20'>
                {/* <thead>
                    <tr>
                        <td className='text-center'>Todo</td>
                        <td className='text-center'>Actions</td>
                    </tr>
                </thead> */}
                <tbody className=''>
                <tr>
                    <td className='w-5/6'>Lorem Ipsum</td>
                    <td className='text-center'>
                        <a href='/done' className='text-green-500 mx-1'><FontAwesomeIcon icon={faCheckSquare} /></a>
                        <a href='/update' className='text-yellow-500 mx-1'><FontAwesomeIcon icon={faPenSquare} /></a>
                        <a href='/delete' className='text-red-500 mx-1'><FontAwesomeIcon icon={faXmarkSquare} /></a>
                    </td>
                </tr>
                </tbody>
            </table>
        </section>
    );
}

export default ToDoList;