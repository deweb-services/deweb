---
sidebar_position: 1
---

# Set up an SSH key

When you set up SSH key, you create a key pair that contains a private key (saved to your local computer) and a public key (uploaded to Deployment VPS). VPS uses the key pair to authenticate anything the associated account can access. This two-way mechanism prevents man-in-the-middle attacks.

This first key pair is your default SSH identity.

:::info

For security reasons, we recommend that you generate a new SSH key and replace the existing key on your account at least once a year.

:::

## Set up SSH for Git on Windows

Use this section to create a default identity and SSH key when you're using Git on Windows. By default, the system adds keys for all identities to the `/Users/<username>/.ssh` directory.

### Step 1. Set up your default identity

1. From the command line, enter ssh-keygen.

:::info For Windows 7 or earlier

You can only enter `ssh-keygen` into the Git Bash window. It won't work in the Command prompt.

:::

The command prompts you for a file to save the key in:

```bash
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/c/Users/emmap1/.ssh/id_rsa):
```

2. Press enter to accept the default key and path, `/c/Users/<username>/.ssh/id_rsa`.

:::info

We recommend keeping the default key name unless you have a reason to change it. To create a key with a non-default name or path, specify the full path to the key. For example, to create a key called `my-new-ssh-key`, enter the Windows path, shown here:

```
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/c/Users/emmap1/.ssh/id_rsa): c:\Users\emmap1\.ssh\my-new-ssh-key
```

:::

3. Enter and re-enter a passphrase when prompted.

The command creates your default identity with its public and private keys. The whole interaction looks similar to this:

```bash
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/c/Users/emmap1/.ssh/id_rsa):
Created directory '/c/Users/emmap1/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /c/Users/emmap1/.ssh/id_rsa.
Your public key has been saved in /c/Users/emmap1/.ssh/id_rsa.pub.
The key fingerprint is: e7:94:d1:a3:02:ee:38:6e:a4:5e:26:a3:a9:f4:95:d4 emmap1@EMMA-PC
```

4. List the contents of .ssh to view the key files.

You should see something like the following:

```bash
$ dir .ssh
id_rsa  id_rsa.pub
```

The command displays two files, one for the public key (for example `id_rsa.pub`) and one for the private key (for example, `id_rsa`).

### Step 2. Add the key to the ssh-agent

If you don't want to type your password each time you use the key, you'll need to add it to the ssh-agent.

1. To start the agent, run the following:

```bash
$ eval $(ssh-agent)
Agent pid 9700
```

2. Enter ssh-add followed by the path to the private key file:

```bash
$ ssh-add ~/.ssh/<private_key_file>
```

### Step 3. Get your public key

Open your .ssh/id_rsa.pub file (or whatever you named the public key file) and copy its contents.
You may see an email address on the last line. It doesn't matter whether or not you include the email address.

## Set up SSH on macOS/Linux

Use this section to create a default identity and SSH key on macOS or Linux. By default, the system adds keys to the `/Users/<yourname>/.ssh` directory on macOS and `/home/<username>/.ssh` on Linux.

### Step 1. Set up your default identity

1. From the terminal, enter `ssh-keygen` at the command line.
   The command prompts you for a file to save the key in:

```bash
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/emmap1/.ssh/id_rsa):
```

2. Press the Enter or Return key to accept the default location.

:::info

We recommend you keep the default key name unless you have a reason to change it.

To create a key with a name or path other than the default, specify the full path to the key. For example, to create a key called `my-new-ssh-key`, enter a path like the one shown at the prompt:

```bash
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/emmap1/.ssh/id_rsa): /Users/emmap1/.ssh/my-new-ssh-key
```

:::

3. Enter and re-enter a passphrase when prompted.
   The command creates your default identity with its public and private keys. The whole interaction will look similar to the following:

```bash
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/emmap1/.ssh/id_rsa):
Created directory '/Users/emmap1/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /Users/emmap1/.ssh/id_rsa.
Your public key has been saved in /Users/emmap1/.ssh/id_rsa.pub.
The key fingerprint is:
4c:80:61:2c:00:3f:9d:dc:08:41:2e:c0:cf:b9:17:69 emmap1@myhost.local
The key's randomart image is:
+--[ RSA 2048]----+
|*o+ooo.          |
|.+.=o+ .         |
|. *.* o .        |
| . = E o         |
|    o . S        |
|   . .           |
|     .           |
|                 |
|                 |
+-----------------+
```

4. List the contents of ~/.ssh to view the key files.

```bash
$ ls ~/.ssh
id_rsa  id_rsa.pub
```

The command displays two files, one for the public key (for example `id_rsa.pub`) and one for the private key (for example, `id_rsa`).

### Step 2. Get your public key

In your terminal window, copy the contents of your public key file. If you renamed the key, replace `id_rsa.pub` with the public key file name.

On Linux, you can cat the contents:

```bash
$ cat ~/.ssh/id_rsa.pub
```

On macOS, the following command copies the output to the clipboard:

```bash
$ pbcopy < ~/.ssh/id_rsa.pub
```
