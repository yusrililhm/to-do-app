function Header() {
    return (
        <header className="flex justify-between items-center fixed left-0 right-0 bg-white shadow-md h-16 font-semibold font-Poppins">
            <nav>
                <ul className="flex items-center justify-center mx-12">
                    <li><a href="/">Todo</a></li>
                </ul>
            </nav>
            <nav>
                <ul className="flex items-center justify-center mx-12">
                    <li className="mx-4 text-blue-900"><a href="/">Sign In</a></li>
                    <li className="mx-4 bg-teal-500 text-white py-3 px-2 rounded-md"><a href="/">Getting Started</a></li>
                </ul>
            </nav>
        </header>
    );
}

export default Header;