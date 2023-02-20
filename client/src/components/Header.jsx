import React from "react";
import { Navbar } from "flowbite-react";
import { Link } from "react-router-dom";

export default function Header(props) {
	return (
		<>
			<Navbar {...props} fluid={true} rounded={true}>
				<div className='max-w-screen-xl w-full mx-auto md:flex md:justify-between px-2'>
					<div className='w-full md:w-auto flex md:block justify-between mb-2'>
						<Navbar.Brand as={Link} to='/'>
							<img src='/img/Waysbeans.svg' alt='waysbook Logo' />
						</Navbar.Brand>
						<Navbar.Toggle />
					</div>
					{props.children}
				</div>
			</Navbar>
			<div className='h-20' />
		</>
	);
}
