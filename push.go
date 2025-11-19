package shanghai

import (
	"fmt"
	"os/exec"
	"strings"
)

func PushImages(s *session, i string) error {
	ims := getImages(s, i)

	for _, im := range ims {
		tags := im.Tags()

		if s.dryRun {
			// In dry-run mode, just print what would be done
			for _, tag := range tags {
				fmt.Printf("DRY RUN: pushing image '%s'\n", tag)
			}
			continue
		}

		for _, tag := range tags {
			if strings.HasPrefix(tag, "localhost/") {
				s.l.Out.Write([]byte(fmt.Sprintf("Skipping tag '%s'\n", tag)))
				continue
			}

			if err := pushImage(s.l, tag, s.c.Engine); err != nil {
				return fmt.Errorf("failed to push tag '%s': %w", tag, err)
			}
		}
	}

	return nil
}

func PushGroup(s *session, g string) error {
	names := s.f.Groups[g]
	for _, name := range names {
		if s.dryRun {
			// In dry-run mode, just print what would be done
			fmt.Printf("DRY RUN: pushing group '%s'\n", name)
			if err := PushImages(s, name); err != nil {
				return fmt.Errorf("failed to push image from group '%s' during dry run: %w", g, err)
			}

			continue
		}

		if err := PushImages(s, name); err != nil {
			return fmt.Errorf("failed to push image from group '%s': %w", g, err)
		}
	}

	return nil
}

func pushImage(lw LogWriters, t string, e string) error {
	cmd := exec.Command(e, "push", t)

	cmd.Stderr = lw.Err
	cmd.Stdout = lw.Out

	lw.Out.Write([]byte(fmt.Sprintf("Pushing %s\n", t)))
	if err := cmd.Run(); err != nil {
		lw.Err.Write([]byte(fmt.Sprintf("failed to run push command for '%s': %s\n", t, err.Error())))
		return fmt.Errorf("failed to run push command for '%s': %w", t, err)
	}
	lw.Out.Write([]byte(fmt.Sprintf("Push done for %s\n", t)))

	return nil
}
