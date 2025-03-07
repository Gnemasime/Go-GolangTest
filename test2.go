package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

type ContactsManager struct {
	contacts []Contact
}

func (cm *ContactsManager) addContact(name, phone, email string) {
	cm.contacts = append(cm.contacts, Contact{Name: name, Phone: phone, Email: email})
}

func (cm *ContactsManager) viewContacts() {
	if len(cm.contacts) == 0 {
		fmt.Println("No contacts available.")
		return
	}
	for i, contact := range cm.contacts {
		fmt.Printf("%d. Name: %s, Phone: %s, Email: %s\n", i+1, contact.Name, contact.Phone, contact.Email)
	}
}

func (cm *ContactsManager) updateContact(index int, name, phone, email string) {
	if index >= 0 && index < len(cm.contacts) {
		cm.contacts[index] = Contact{Name: name, Phone: phone, Email: email}
		fmt.Println("Contact updated successfully.")
	} else {
		fmt.Println("Contact not found.")
	}
}

func (cm *ContactsManager) deleteContact(index int) {
	if index >= 0 && index < len(cm.contacts) {
		cm.contacts = append(cm.contacts[:index], cm.contacts[index+1:]...)
		fmt.Println("Contact deleted successfully.")
	} else {
		fmt.Println("Contact not found.")
	}
}

func main() {
	cm := ContactsManager{}

	for {
		fmt.Println("\nContact Management System")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, phone, email string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter phone number: ")
			fmt.Scan(&phone)
			fmt.Print("Enter email: ")
			fmt.Scan(&email)

			cm.addContact(name, phone, email)
			fmt.Println("Contact added successfully.")

		case 2:
			cm.viewContacts()

		case 3:
			var index int
			var name, phone, email string
			cm.viewContacts()
			fmt.Print("Enter the index of the contact to update: ")
			fmt.Scan(&index)
			index-- // Adjusting for 1-based index
			fmt.Print("Enter new name: ")
			fmt.Scan(&name)
			fmt.Print("Enter new phone number: ")
			fmt.Scan(&phone)
			fmt.Print("Enter new email: ")
			fmt.Scan(&email)

			cm.updateContact(index, name, phone, email)

		case 4:
			var index int
			cm.viewContacts()
			fmt.Print("Enter the index of the contact to delete: ")
			fmt.Scan(&index)
			index-- // Adjusting for 1-based index
			cm.deleteContact(index)

		case 5:
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
