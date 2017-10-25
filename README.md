# Technology Learning Studio
 
* October
	* [October 3rd](#2017-10-03) - Initial Post
	* [October 4th](#2017-10-04) - First Day of Ponzu
	* [October 11th](#2017-10-11) - Actually, It's Docker
	* [October 18th](#2017-10-18) - Actually, It's Network Problems
	* [October 25th](#2017-10-25) - Illustrator and Photoshop Project - Hijabi Skateboarder

## Creating a Reservation System

### 2017-10-03

#### Initial Post

This readme acts as a blog for Professor Howison's Technology Learning Studio course at the University of Texas at Austin. I'll document my learning process here.

### 2017-10-04 

#### First Day of Ponzu

Ponzu is a framework for creating a content management system (CMS) with a built-in database and encryption. My goal is to create a reservation system for the devices in the IT Lab. This won't be a user-facing system, but an administrative system for tracking the devices, who has them checked out, when they are due, their associated accessories, and what kind of condition they're in.

The first day of Ponzu, I dealt a lot with the setup of the environment and the dependencies of the CMS. I've never dealt with Docker before, and I alternated between populating the CMS with the appropriate attributes and the corresponding inputs (short text, text block, select) and researching the best way to deploy it while I'm developing it so that IT staff members can populate it and provide feedback.

After that, my goal is to take advantage of the JSON APIs and addons that have been developed for Ponzu so that when a device is checked out by a student, either the student or the IT Lab or both will be warned when the item is due, overdue, or reserved either by email our through Spiceworks tickets. Figuring out how to automate these parts together will be the most difficult and time-consuming part of the process. Deploying it on a server should just be a matter of coordinating with the IT staff and Googling.

### 2017-10-11 

#### Actually, It's Docker

Creating the reservation system was actually pretty easy. There are other implementations outside of the scope of the learning process that I won't be documenting here (like automating reminders to turn items back in to the IT Lab), so I'll be focusing on learning how to use Docker. Docker provides a virtual environment for the app to live in. I think that learning Docker, which is widely used in the development and deployment of applications (DevOps), will provide a lot of insight for me as a designer and amateur coder. I think it's very important to understand the capabilities and limitations of all of the tools and frameworks that go into designing, creating, and releasing digital products.

I've found a series of Lynda.com Docker tutorials, and completing those confidently is my short goal. My long goal is to deploy the checkout system and to create multiple user-space instances, perhaps one for the IT Lab staff and one for students. I also have a goal of figuring out how to deploy the app in the iSchool server rather than using Amazon Web Services or Google Cloud Services.

After completing the first three and a half sections of the five-part Docker tutorial, I have a much better idea of how Docker actually works. Containers are created from images, and those images are operating systems that run virtually. They can be various distributions of Linux and releases thereof, and you can carefully choose which image to run to make sure that it has the programs and commands that you need. You can also create a network from these images that can send data from one to the other.

There is also something called a Dockerfile, which is basically a script for creating or compiling a Docker image. The content management system that I've made makes a Dockerfile, so I need to figure out how to compile that into an image and then run one or more containers of the CMS that can be accessed by the IT Lab staff. I might see if I can require that someone is on the utexas.edu network before they can access the CMS in the first place for increased security.

In addition to learning the concepts and terminology of Docker, I'm also troubleshooting many small issues that arise. I had to uninstall and reinstall a newer version of a program called VirtualBox, which came with Docker's install but apparently comes out of the box outdated. I had some issues with my permissions for my /var/run folder, where the docker daemon runs. I'm trying to view each of these small setbacks as opportunities to get a better sense of how a Unix system functions, from the filesystem to networking.

### 2017-10-18

#### Actually, It's Network Problems

After the regular amount and severity of setbacks I encountered running Docker locally on my computer, I felt confident in building and running the same image on one of the iSchool's virtual servers. The process went smoothly up until the point of building the app from the Dockerfile. The paths specified in the Dockerfile were not matching up with the GOPATH that I had set up on the virtual server. After editing it a couple of times with some trial and error, I managed to reach the part of the build during which packages would be downloaded from GitHub and the Golang standard library.

The process encountered a fatal error and was unable to resolve github.com and golang.org as hosts. After some Googling, I found a fix to uncomment a line in the Docker upstart configuration file, which would allow Docker to use Google's DNS servers during the build. I made sure to restart Docker so that this configuration would be implemented, and ran the build command again. Same problem!

I double-checked the configuration file, and everything should be good to go. I messaged a friend of mine, a software developer who regularly works with Docker. The good news was that he had encountered this problem many times before and had solved it the same way—utilizing Google's DNS servers. The bad news was that he had no idea why it hadn't worked on my machine. I worked to troubleshoot this problem, but couldn't find any other users that had implemented this fix without it having solved the problem.

I found the Lynda tutorial and the documentation for Docker to be very good. The instructor explained each step very thoroughly and pointed out and emphasized steps that were counterintuitive; for example, creating containers from images and what is carried over from the images and what is not when multiple containers are created. Basically the image is a source, and the containers are representations of that image that can be changed, preserved, or done away with when no longer useful.

Once I encountered network problems, I became very frustrated because I was being sidetracked by external issues that had nothing to do with Docker per se. I suspect that the fact that I'm trying to launch an application on a virtual server in a virtual environment might have have something to do with the network problems that I'm having. Because the ultimate, external goal of learning Docker is to create a reservation system for the IT Lab, the logical thing to do would be to get help from the IT technicians that I work with. I feel that this course of action is outside the scope of my learning plan, yet I'm stuck without it—I've gotten a good feel for Docker and have created and built containers from images. Now I have network problems, and I'm not particularly interested in going this in-depth into DEVOPS. As I stated earlier, I hoped to get a better sense of the environment in which Docker is used, its limitations, and its capabilities. Now that my issues seem to go beyond the working knowledge of a professional software developer that works with Docker on a daily basis, I think I have gone a bit too far into the deep end.

For that reason, I hope to change course to another skill that I know will be useful to me professionally (and recreationally): graphic design. Sordid attempts at designing flyers and promotional materials in the past have been personally embarrassing for me. I use Adobe XD and know it pretty well, but Photoshop and Illustrator are both a little scary for me. I'd like to become more comfortable with both of them, and at the same time develop a more... let's say marketable aesthetic.

### 2017-1025

#### Illustrator and Photoshop Project - Hijabi Skateboarder

I learned a great deal about the basic workflow of both Illustrator and Photoshop by editing a photo in Photoshop by cropping the background out of a photo, creating an image of the person I isolated by making it black and white and playing with the contrast and brightness, then opening it in Illustrator. From there, I created a palette of four colors that I liked, filled in the background, then created some visually pleasing shapes to highlight the person—an Afghani girl doing a kick-flip on her skateboard. I'm very pleased with the final result.

![Afghani Skateboarder](src="https://biglin.io/img/illustrator/afghanigirl.png")

To experiment further, I also made an avocado with some yellow, green, and brown gradients, and a small logo for a fake app that incorporated a leaf and roots into it.
