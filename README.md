# Technology Learning Studio
 
* October
	* [October 3rd](#2017-10-03) - Initial Post
	* [October 4th](#2017-10-04) - First Day of Ponzu
	* [October 11th](#2017-10-11) - Actually, It's Docker

## Creating a Reservation System

### 2017-10-03

#### Initial Post

This readme acts as a blog for Professor Howison's Technology Learning Studio course at the University of Texas at Austin. I'll document my learning process here.

### 2017-10-04 

#### First Day of Ponzu

Ponzu is a framework for creating a content management system (CMS) with a built-in database and encryption. My goal is to create a reservation system for the devices in the IT Lab. This won't be a user-facing system, but an administrative system for tracking the devices, who has them checked out, when they are due, their associated accessories, and what kind of condition they're in.

The first day of Ponzu, I dealt a lot with the setup of the environment and the dependencies of the CMS. I've never dealt with Docker before, and I alternated between populating the CMS with the appropriate attributes and the corresponding inputs (short text, text block, select) and researching the best way to deploy it while I'm developing it so that IT staff members can populate it and provide feedback.

After that, my goal is to take advantage of the JSON APIs and addons that have been developed for Ponzu so that when a device is checked out by a student, either the student or the IT Lab or both will be warned when the item is due, overdue, or reserved either by email our through Spiceworks tickets. Figuring out how to automate these parts together will be the most difficult and time-consuming part of the process. Deploying it on a server should just be a matter of coordinating with the IT staff and Googling.

#### Actually, It's Docker

Creating the reservation system was actually pretty easy. There are other implementations outside of the scope of the learning process that I won't be documenting here (like automating reminders to turn items back in to the IT Lab), so I'll be focusing on learning how to use Docker. Docker provides a virtual environment for the app to live in. I think that learning Docker, which is widely used in the development and deployment of applications (DevOps), will provide a lot of insight for me as a designer and amateur coder. I think it's very important to understand the capabilities and limitations of all of the tools and frameworks that go into designing, creating, and releasing digital products.

I've found a series of Lynda.com Docker tutorials, and completing those confidently is my short goal. My long goal is to deploy the checkout system and to create multiple user-space instances, perhaps one for the IT Lab staff and one for students. I also have a goal of figuring out how to deploy the app in the iSchool server rather than using Amazon Web Services or Google Cloud Services.

After completing the first three and a half sections of the five-part Docker tutorial, I have a much better idea of how Docker actually works. Containers are created from images, and those images are operating systems that run virtually. They can be various distributions of Linux and releases thereof, and you can carefully choose which image to run to make sure that it has the programs and commands that you need. You can also create a network from these images that can send data from one to the other.

There is also something called a Dockerfile, which is basically a script for creating or compiling a Docker image. The content management system that I've made makes a Dockerfile, so I need to figure out how to compile that into an image and then run one or more containers of the CMS that can be accessed by the IT Lab staff. I might see if I can require that someone is on the utexas.edu network before they can access the CMS in the first place for increased security.

In addition to learning the concepts and terminology of Docker, I'm also troubleshooting many small issues that arise. I had to uninstall and reinstall a newer version of a program called VirtualBox, which came with Docker's install but apparently comes out of the box outdated. I had some issues with my permissions for my /var/run folder, where the docker daemon runs. I'm trying to view each of these small setbacks as opportunities to get a better sense of how a Unix system functions, from the filesystem to networking.
