import React from 'react';

export function Footer() {
    const currentYear = new Date().getFullYear()
    return(
        <div>
            <hr/>
            <p className="footer">Copyright &copy; {currentYear} ProjectHub</p>
        </div>
    );
}