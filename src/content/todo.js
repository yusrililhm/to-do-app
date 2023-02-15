import '../assets/css/index.css'

function Todo() {
    return (
        <section className='font-Poppins h-screen'>
            <div className='flex justify-center items-center'>
                <form method='POST' action='/Add' className='flex justify-center items-center mt-20 bg-white shadow-lg py-2 px-4 rounded-full w-3/5'>
                    <input type='text' name='todo' id='todo' placeholder='Type your plan here' className='w-full'/>
                </form>
            </div>
        </section>
    );
}

export default Todo;