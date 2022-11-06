"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[5],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>c});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},s=Object.keys(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var p=a.createContext({}),l=function(e){var t=a.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},u=function(e){var t=l(e.components);return a.createElement(p.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},h=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,s=e.originalType,p=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),h=l(n),c=r,y=h["".concat(p,".").concat(c)]||h[c]||d[c]||s;return n?a.createElement(y,o(o({ref:t},u),{},{components:n})):a.createElement(y,o({ref:t},u))}));function c(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var s=n.length,o=new Array(s);o[0]=h;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i.mdxType="string"==typeof e?e:r,o[1]=i;for(var l=2;l<s;l++)o[l]=n[l];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}h.displayName="MDXCreateElement"},7152:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>o,default:()=>d,frontMatter:()=>s,metadata:()=>i,toc:()=>l});var a=n(7462),r=(n(7294),n(3905));const s={sidebar_position:1},o="Set up an SSH key",i={unversionedId:"support/deployments/set-up-an-ssh-key",id:"support/deployments/set-up-an-ssh-key",title:"Set up an SSH key",description:"When you set up SSH key, you create a key pair that contains a private key (saved to your local computer) and a public key (uploaded to Deployment VPS). VPS uses the key pair to authenticate anything the associated account can access. This two-way mechanism prevents man-in-the-middle attacks.",source:"@site/docs/support/deployments/set-up-an-ssh-key.md",sourceDirName:"support/deployments",slug:"/support/deployments/set-up-an-ssh-key",permalink:"/support/deployments/set-up-an-ssh-key",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"defaultSidebar",previous:{title:"DWS DNS Server",permalink:"/domains/dws-dns-server"}},p={},l=[{value:"Set up SSH for Git on Windows",id:"set-up-ssh-for-git-on-windows",level:2},{value:"Step 1. Set up your default identity",id:"step-1-set-up-your-default-identity",level:3},{value:"Step 2. Add the key to the ssh-agent",id:"step-2-add-the-key-to-the-ssh-agent",level:3},{value:"Step 3. Get your public key",id:"step-3-get-your-public-key",level:3},{value:"Set up SSH on macOS/Linux",id:"set-up-ssh-on-macoslinux",level:2},{value:"Step 1. Set up your default identity",id:"step-1-set-up-your-default-identity-1",level:3},{value:"Step 2. Get your public key",id:"step-2-get-your-public-key",level:3}],u={toc:l};function d(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,a.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"set-up-an-ssh-key"},"Set up an SSH key"),(0,r.kt)("p",null,"When you set up SSH key, you create a key pair that contains a private key (saved to your local computer) and a public key (uploaded to Deployment VPS). VPS uses the key pair to authenticate anything the associated account can access. This two-way mechanism prevents man-in-the-middle attacks."),(0,r.kt)("p",null,"This first key pair is your default SSH identity."),(0,r.kt)("admonition",{type:"info"},(0,r.kt)("p",{parentName:"admonition"},"For security reasons, we recommend that you generate a new SSH key and replace the existing key on your account at least once a year.")),(0,r.kt)("h2",{id:"set-up-ssh-for-git-on-windows"},"Set up SSH for Git on Windows"),(0,r.kt)("p",null,"Use this section to create a default identity and SSH key when you're using Git on Windows. By default, the system adds keys for all identities to the ",(0,r.kt)("inlineCode",{parentName:"p"},"/Users/<username>/.ssh")," directory."),(0,r.kt)("h3",{id:"step-1-set-up-your-default-identity"},"Step 1. Set up your default identity"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"From the command line, enter ssh-keygen.")),(0,r.kt)("admonition",{title:"For Windows 7 or earlier",type:"info"},(0,r.kt)("p",{parentName:"admonition"},"You can only enter ",(0,r.kt)("inlineCode",{parentName:"p"},"ssh-keygen")," into the Git Bash window. It won't work in the Command prompt.")),(0,r.kt)("p",null,"The command prompts you for a file to save the key in:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ssh-keygen\nGenerating public/private rsa key pair.\nEnter file in which to save the key (/c/Users/emmap1/.ssh/id_rsa):\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"Press enter to accept the default key and path, ",(0,r.kt)("inlineCode",{parentName:"li"},"/c/Users/<username>/.ssh/id_rsa"),".")),(0,r.kt)("admonition",{type:"info"},(0,r.kt)("p",{parentName:"admonition"},"We recommend keeping the default key name unless you have a reason to change it. To create a key with a non-default name or path, specify the full path to the key. For example, to create a key called ",(0,r.kt)("inlineCode",{parentName:"p"},"my-new-ssh-key"),", enter the Windows path, shown here:"),(0,r.kt)("pre",{parentName:"admonition"},(0,r.kt)("code",{parentName:"pre"},"$ ssh-keygen\nGenerating public/private rsa key pair.\nEnter file in which to save the key (/c/Users/emmap1/.ssh/id_rsa): c:\\Users\\emmap1\\.ssh\\my-new-ssh-key\n"))),(0,r.kt)("ol",{start:3},(0,r.kt)("li",{parentName:"ol"},"Enter and re-enter a passphrase when prompted.")),(0,r.kt)("p",null,"The command creates your default identity with its public and private keys. The whole interaction looks similar to this:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ssh-keygen\nGenerating public/private rsa key pair.\nEnter file in which to save the key (/c/Users/emmap1/.ssh/id_rsa):\nCreated directory '/c/Users/emmap1/.ssh'.\nEnter passphrase (empty for no passphrase):\nEnter same passphrase again:\nYour identification has been saved in /c/Users/emmap1/.ssh/id_rsa.\nYour public key has been saved in /c/Users/emmap1/.ssh/id_rsa.pub.\nThe key fingerprint is: e7:94:d1:a3:02:ee:38:6e:a4:5e:26:a3:a9:f4:95:d4 emmap1@EMMA-PC\n")),(0,r.kt)("ol",{start:4},(0,r.kt)("li",{parentName:"ol"},"List the contents of .ssh to view the key files.")),(0,r.kt)("p",null,"You should see something like the following:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ dir .ssh\nid_rsa  id_rsa.pub\n")),(0,r.kt)("p",null,"The command displays two files, one for the public key (for example ",(0,r.kt)("inlineCode",{parentName:"p"},"id_rsa.pub"),") and one for the private key (for example, ",(0,r.kt)("inlineCode",{parentName:"p"},"id_rsa"),")."),(0,r.kt)("h3",{id:"step-2-add-the-key-to-the-ssh-agent"},"Step 2. Add the key to the ssh-agent"),(0,r.kt)("p",null,"If you don't want to type your password each time you use the key, you'll need to add it to the ssh-agent."),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"To start the agent, run the following:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ eval $(ssh-agent)\nAgent pid 9700\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"Enter ssh-add followed by the path to the private key file:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ssh-add ~/.ssh/<private_key_file>\n")),(0,r.kt)("h3",{id:"step-3-get-your-public-key"},"Step 3. Get your public key"),(0,r.kt)("p",null,"Open your .ssh/id_rsa.pub file (or whatever you named the public key file) and copy its contents.\nYou may see an email address on the last line. It doesn't matter whether or not you include the email address."),(0,r.kt)("h2",{id:"set-up-ssh-on-macoslinux"},"Set up SSH on macOS/Linux"),(0,r.kt)("p",null,"Use this section to create a default identity and SSH key on macOS or Linux. By default, the system adds keys to the ",(0,r.kt)("inlineCode",{parentName:"p"},"/Users/<yourname>/.ssh")," directory on macOS and ",(0,r.kt)("inlineCode",{parentName:"p"},"/home/<username>/.ssh")," on Linux."),(0,r.kt)("h3",{id:"step-1-set-up-your-default-identity-1"},"Step 1. Set up your default identity"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"From the terminal, enter ",(0,r.kt)("inlineCode",{parentName:"li"},"ssh-keygen")," at the command line.\nThe command prompts you for a file to save the key in:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ssh-keygen\nGenerating public/private rsa key pair.\nEnter file in which to save the key (/Users/emmap1/.ssh/id_rsa):\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"Press the Enter or Return key to accept the default location.")),(0,r.kt)("admonition",{type:"info"},(0,r.kt)("p",{parentName:"admonition"},"We recommend you keep the default key name unless you have a reason to change it."),(0,r.kt)("p",{parentName:"admonition"},"To create a key with a name or path other than the default, specify the full path to the key. For example, to create a key called ",(0,r.kt)("inlineCode",{parentName:"p"},"my-new-ssh-key"),", enter a path like the one shown at the prompt:"),(0,r.kt)("pre",{parentName:"admonition"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ssh-keygen\nGenerating public/private rsa key pair.\nEnter file in which to save the key (/Users/emmap1/.ssh/id_rsa): /Users/emmap1/.ssh/my-new-ssh-key\n"))),(0,r.kt)("ol",{start:3},(0,r.kt)("li",{parentName:"ol"},"Enter and re-enter a passphrase when prompted.\nThe command creates your default identity with its public and private keys. The whole interaction will look similar to the following:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ssh-keygen\nGenerating public/private rsa key pair.\nEnter file in which to save the key (/Users/emmap1/.ssh/id_rsa):\nCreated directory '/Users/emmap1/.ssh'.\nEnter passphrase (empty for no passphrase):\nEnter same passphrase again:\nYour identification has been saved in /Users/emmap1/.ssh/id_rsa.\nYour public key has been saved in /Users/emmap1/.ssh/id_rsa.pub.\nThe key fingerprint is:\n4c:80:61:2c:00:3f:9d:dc:08:41:2e:c0:cf:b9:17:69 emmap1@myhost.local\nThe key's randomart image is:\n+--[ RSA 2048]----+\n|*o+ooo.          |\n|.+.=o+ .         |\n|. *.* o .        |\n| . = E o         |\n|    o . S        |\n|   . .           |\n|     .           |\n|                 |\n|                 |\n+-----------------+\n")),(0,r.kt)("ol",{start:4},(0,r.kt)("li",{parentName:"ol"},"List the contents of ~/.ssh to view the key files.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ ls ~/.ssh\nid_rsa  id_rsa.pub\n")),(0,r.kt)("p",null,"The command displays two files, one for the public key (for example ",(0,r.kt)("inlineCode",{parentName:"p"},"id_rsa.pub"),") and one for the private key (for example, ",(0,r.kt)("inlineCode",{parentName:"p"},"id_rsa"),")."),(0,r.kt)("h3",{id:"step-2-get-your-public-key"},"Step 2. Get your public key"),(0,r.kt)("p",null,"In your terminal window, copy the contents of your public key file. If you renamed the key, replace ",(0,r.kt)("inlineCode",{parentName:"p"},"id_rsa.pub")," with the public key file name."),(0,r.kt)("p",null,"On Linux, you can cat the contents:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ cat ~/.ssh/id_rsa.pub\n")),(0,r.kt)("p",null,"On macOS, the following command copies the output to the clipboard:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"$ pbcopy < ~/.ssh/id_rsa.pub\n")))}d.isMDXComponent=!0}}]);